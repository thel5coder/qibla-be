package usecase

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"qibla-backend/helpers/enums"
	"qibla-backend/helpers/messages"
	"qibla-backend/helpers/str"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type FaspayUseCase struct {
	*UcContract
}

func (uc FaspayUseCase) getSignature(compose string) (res string) {
	var md5 = md5.New()
	var sha1 = sha1.New()
	md5.Write([]byte(compose))
	md5EncryptedStr := md5.Sum(nil)
	md5Str := fmt.Sprintf("%x", md5EncryptedStr)

	sha1.Write([]byte(md5Str))
	sha1EncryptedStr := sha1.Sum(nil)
	res = fmt.Sprintf("%x", sha1EncryptedStr)

	return res
}

func (uc FaspayUseCase) GetLisPaymentMethods() (res map[string]interface{}, err error) {
	compose := os.Getenv("FASPAY_USER_ID") + `` + os.Getenv("FASPAY_PASSWORD")
	signature := uc.getSignature(compose)
	var client http.Client
	var bodyPost = []byte(`{"merchant_id":"` + os.Getenv("FASPAY_MERCHANT_ID") + `","signature":"` + signature + `"}`)

	request, err := http.NewRequest("POST", fasPayBaseUrl+"/100001/10", bytes.NewBuffer(bodyPost))
	if err != nil {
		fmt.Print(err.Error())
		return res, err
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		return res, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		return res, err
	}
	err = json.Unmarshal(body, &res)

	return res, err
}

func (uc FaspayUseCase) PostData(input requests.FaspayPostRequest, contact viewmodel.ContactVm) (res map[string]interface{}, err error) {
	compose := os.Getenv("FASPAY_USER_ID") + `` + os.Getenv("FASPAY_PASSWORD") + `` + input.InvoiceNumber
	phoneNumber := str.StringToInt(contact.PhoneNumber)
	signature := uc.getSignature(compose)
	var client http.Client
	var items []viewmodel.ItemFaspayPostDataVm

	for _, item := range input.Item{
		items = append(items,viewmodel.ItemFaspayPostDataVm{
			Product:     item.Product,
			Amount:      item.Amount,
			Qty:         item.Qty,
			PaymentPlan: item.PaymentPlan,
			Tenor:       item.Tenor,
			MerchantID:  os.Getenv("FASPAY_MERCHANT_ID"),
		})
	}
	body := viewmodel.FaspayPostDataVm{
		Request:        input.RequestTransaction,
		MerchantID:     os.Getenv("FASPAY_MERCHANT_ID"),
		Merchant:       contact.TravelAgentName,
		BillNo:         input.InvoiceNumber,
		BillDate:       input.TransactionDate,
		BillExpired:    input.DueDate,
		BillDesc:       input.TransactionDesc,
		BillCurrency:   defaultFaspayCurrency,
		BillTotal:      input.Total,
		PaymentChannel: input.PaymentChannel,
		PayType:        defaultFaspayPayType,
		CustNo:         input.UserID,
		CustName:       contact.TravelAgentName,
		Msisdn:         phoneNumber,
		Email:          contact.Email,
		Terminal:       defaultFaspayTerminal,
		Signature:      signature,
		Item:           items,
	}
	bodyPost, _ := json.Marshal(body)


	request, err := http.NewRequest("POST", fasPayBaseUrl+"/300011/10", bytes.NewBuffer(bodyPost))
	if err != nil {
		return res, err
	}
	response, err := client.Do(request)
	if err != nil {
		return res, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		return res, err
	}
	fmt.Println(res)
	if res["response_code"] != "00"{
		return res,errors.New(messages.PaymentFailed)
	}

	return res, err
}

func (uc FaspayUseCase) CheckPaymentStatus(invoiceID string) (res map[string]interface{},err error){
	var client http.Client
	transactionUc := TransactionUseCase{UcContract:uc.UcContract}
	transaction,err := transactionUc.ReadBy("t.id",invoiceID,"=")
	if err != nil {
		return res,err
	}
	signature := uc.getSignature(os.Getenv("FASPAY_USER_ID")+``+os.Getenv("FASPAY_PASSWORD")+``+transaction.InvoiceNumber)
	body := viewmodel.CheckPaymentStatus{
		Request:    enums.KeyTransactionType6,
		TrxID:      transaction.TrxID,
		MerchantID: os.Getenv("FASPAY_MERCHANT_ID"),
		BillNo:     transaction.InvoiceNumber,
		Signature:  signature,
	}

	bodyPost,_ := json.Marshal(body)
	request, err := http.NewRequest("POST", fasPayBaseUrl+"/100004/10", bytes.NewBuffer(bodyPost))
	if err != nil {
		return res, err
	}
	response, err := client.Do(request)
	if err != nil {
		return res, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		return res, err
	}

	return res,err
}

func (uc FaspayUseCase) PaymentNotification(input *requests.PaymentNotificationRequest) (res viewmodel.PaymentNotificationVm,err error,code int){
	fmt.Println("uc")
	fmt.Println(input.PaymentDate)
	now := time.Now().UTC().Format("2006-01-02 15:04:05")
	transactionUc := TransactionUseCase{UcContract:uc.UcContract}
	transaction,err := transactionUc.ReadBy("t.trx_id",input.TrxID,"=")
	if err != nil {
		fmt.Println(1)

		res = viewmodel.PaymentNotificationVm{
			Response:     "Payment Notification",
			TrxID:        input.TrxID,
			MerchantID:   input.MerchantID,
			Merchant:     input.Merchant,
			BillNo:       input.BillNo,
			ResponseCode: "05",
			ResponseDesc: "Tagihan tidak ditemukan",
			ResponseDate: now,
		}
		return res,err,http.StatusUnprocessableEntity
	}
	signature := uc.getSignature(os.Getenv("FASPAY_USER_ID")+``+os.Getenv("FASPAY_PASSWORD")+``+transaction.InvoiceNumber+`2`)
	if signature != input.Signature {
		fmt.Println(2)

		res = viewmodel.PaymentNotificationVm{
			Response:     "Payment Notification",
			TrxID:        input.TrxID,
			MerchantID:   input.MerchantID,
			Merchant:     input.Merchant,
			BillNo:       input.BillNo,
			ResponseCode: "09",
			ResponseDesc: "Unknown",
			ResponseDate: now,
		}
		return res,err,http.StatusUnauthorized
	}

	uc.TX,err = uc.DB.Begin()
	if err != nil {
		fmt.Println(3)

		uc.TX.Rollback()
		res = viewmodel.PaymentNotificationVm{
			Response:     "Payment Notification",
			TrxID:        input.TrxID,
			MerchantID:   input.MerchantID,
			Merchant:     input.Merchant,
			BillNo:       input.BillNo,
			ResponseCode: "09",
			ResponseDesc: "Unknown",
			ResponseDate: now,
		}

		return res,err,http.StatusUnprocessableEntity
	}

	transactionUc.TX = uc.TX
	err = transactionUc.EditStatus(transaction.ID,enums.KeyPaymentStatus3,input.PaymentDate)
	if err != nil {
		fmt.Println(err.Error())
		uc.TX.Rollback()
		res = viewmodel.PaymentNotificationVm{
			Response:     "Payment Notification",
			TrxID:        input.TrxID,
			MerchantID:   input.MerchantID,
			Merchant:     input.Merchant,
			BillNo:       input.BillNo,
			ResponseCode: "00",
			ResponseDesc: "Sukses",
			ResponseDate: now,
		}

		return res,err,http.StatusOK
	}

	transactionHistoryUc := TransactionHistoryUseCase{UcContract:uc.UcContract}
	err = transactionHistoryUc.Add(input.TrxID,enums.KeyPaymentStatus3,map[string]interface{}{})
	if err != nil {
		fmt.Println(5)

		uc.TX.Rollback()
		res = viewmodel.PaymentNotificationVm{
			Response:     "Payment Notification",
			TrxID:        input.TrxID,
			MerchantID:   input.MerchantID,
			Merchant:     input.Merchant,
			BillNo:       input.BillNo,
			ResponseCode: "00",
			ResponseDesc: "Sukses",
			ResponseDate: now,
		}

		return res,err,http.StatusOK
	}

	res = viewmodel.PaymentNotificationVm{
		Response:     "Payment Notification",
		TrxID:        input.TrxID,
		MerchantID:   input.MerchantID,
		Merchant:     input.Merchant,
		BillNo:       input.BillNo,
		ResponseCode: "00",
		ResponseDesc: "Sukses",
		ResponseDate: now,
	}
	uc.TX.Commit()

	return res,err,http.StatusOK
}
