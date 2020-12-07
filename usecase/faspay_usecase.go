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
	"qibla-backend/pkg/enums"
	"qibla-backend/pkg/interfacepkg"
	"qibla-backend/pkg/messages"
	"qibla-backend/pkg/str"
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

	request, err := http.NewRequest("POST", os.Getenv("FASPAY_BASE_URL")+"/100001/10", bytes.NewBuffer(bodyPost))
	if err != nil {
		return res, err
	}

	response, err := client.Do(request)
	if err != nil {
		return res, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(body, &res)

	return res, err
}

func (uc FaspayUseCase) GetLisPaymentMethodsByCode(code string) (res string, err error) {
	paymentMethod, err := uc.GetLisPaymentMethods()
	if err != nil {
		return res, err
	}

	var data viewmodel.FaspayPaymentMethodVM
	interfacepkg.Convert(paymentMethod, &data)

	for _, d := range data.PaymentChannel {
		if d.PgCode == code {
			res = d.PgName
		}
	}

	if res == "" {
		return res, errors.New("Empty Data")
	}

	return res, err

}

func (uc FaspayUseCase) PostData(input requests.FaspayPostRequest) (res map[string]interface{}, err error) {
	compose := os.Getenv("FASPAY_USER_ID") + `` + os.Getenv("FASPAY_PASSWORD") + `` + input.InvoiceNumber
	signature := uc.getSignature(compose)
	var client http.Client
	var items []viewmodel.ItemFaspayPostDataVm

	for _, item := range input.Item {
		items = append(items, viewmodel.ItemFaspayPostDataVm{
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
		Merchant:       os.Getenv("FASPAY_MERCHANT"),
		BillNo:         input.InvoiceNumber,
		BillDate:       input.TransactionDate,
		BillExpired:    input.DueDate,
		BillDesc:       input.TransactionDesc,
		BillCurrency:   defaultFaspayCurrency,
		BillTotal:      input.Total,
		PaymentChannel: input.PaymentChannel,
		PayType:        defaultFaspayPayType,
		CustNo:         input.UserID,
		CustName:       input.CustomerName,
		Msisdn:         str.StringToInt(input.CustomerPhoneNumber),
		Email:          input.CustomerEmail,
		Terminal:       defaultFaspayTerminal,
		Signature:      signature,
		Item:           items,
	}
	bodyPost, _ := json.Marshal(body)

	request, err := http.NewRequest("POST", os.Getenv("FASPAY_BASE_URL")+"/300011/10", bytes.NewBuffer(bodyPost))
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
	if res["response_code"] != "00" {
		return res, errors.New(messages.PaymentFailed)
	}

	return res, err
}

func (uc FaspayUseCase) CheckPaymentStatus(invoiceID string) (res map[string]interface{}, err error) {
	var client http.Client
	transactionUc := TransactionUseCase{UcContract: uc.UcContract}
	transaction, err := transactionUc.ReadBy("t.id", invoiceID, "=")
	if err != nil {
		return res, err
	}
	signature := uc.getSignature(os.Getenv("FASPAY_USER_ID") + `` + os.Getenv("FASPAY_PASSWORD") + `` + transaction.InvoiceNumber)
	body := viewmodel.CheckPaymentStatus{
		Request:    enums.KeyTransactionType6,
		TrxID:      transaction.TrxID,
		MerchantID: os.Getenv("FASPAY_MERCHANT_ID"),
		BillNo:     transaction.InvoiceNumber,
		Signature:  signature,
	}

	bodyPost, _ := json.Marshal(body)
	request, err := http.NewRequest("POST", os.Getenv("FASPAY_BASE_URL")+"/100004/10", bytes.NewBuffer(bodyPost))
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

	return res, err
}

func (uc FaspayUseCase) PaymentNotification(input *requests.PaymentNotificationRequest) (res viewmodel.PaymentNotificationVm, err error, code int) {
	now := time.Now().UTC().Format("2006-01-02 15:04:05")

	//check is transaction exist by trx_id
	transactionUc := TransactionUseCase{UcContract: uc.UcContract}
	transaction, err := transactionUc.ReadBy("t.trx_id", input.TrxID, "=")
	if err != nil {
		res = viewmodel.PaymentNotificationVm{
			Response:      "Payment Notification",
			TransactionID: transaction.ID,
			TrxID:         input.TrxID,
			MerchantID:    input.MerchantID,
			Merchant:      input.Merchant,
			BillNo:        input.BillNo,
			ResponseCode:  "05",
			ResponseDesc:  "Tagihan tidak ditemukan",
			ResponseDate:  now,
		}
		return res, err, http.StatusUnprocessableEntity
	}

	//check is signature valid
	signature := uc.getSignature(os.Getenv("FASPAY_USER_ID") + `` + os.Getenv("FASPAY_PASSWORD") + `` + transaction.InvoiceNumber + `2`)
	if signature != input.Signature {
		res = viewmodel.PaymentNotificationVm{
			Response:      "Payment Notification",
			TransactionID: transaction.ID,
			TrxID:         input.TrxID,
			MerchantID:    input.MerchantID,
			Merchant:      input.Merchant,
			BillNo:        input.BillNo,
			ResponseCode:  "09",
			ResponseDesc:  "Unknown",
			ResponseDate:  now,
		}
		return res, err, http.StatusUnauthorized
	}

	//init transaction
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()
		res = viewmodel.PaymentNotificationVm{
			Response:      "Payment Notification",
			TransactionID: transaction.ID,
			TrxID:         input.TrxID,
			MerchantID:    input.MerchantID,
			Merchant:      input.Merchant,
			BillNo:        input.BillNo,
			ResponseCode:  "09",
			ResponseDesc:  "Unknown",
			ResponseDate:  now,
		}

		return res, err, http.StatusUnprocessableEntity
	}

	//edit payment status in transaction
	transactionUc.TX = uc.TX
	err = transactionUc.EditStatus(transaction.ID, enums.KeyPaymentStatus3, input.PaymentDate)
	if err != nil {
		uc.TX.Rollback()
		res = viewmodel.PaymentNotificationVm{
			Response:      "Payment Notification",
			TransactionID: transaction.ID,
			TrxID:         input.TrxID,
			MerchantID:    input.MerchantID,
			Merchant:      input.Merchant,
			BillNo:        input.BillNo,
			ResponseCode:  "00",
			ResponseDesc:  "Sukses",
			ResponseDate:  now,
		}

		return res, err, http.StatusOK
	}

	//add transaction history
	transactionHistoryUc := TransactionHistoryUseCase{UcContract: uc.UcContract}
	err = transactionHistoryUc.Add(input.TrxID, enums.KeyPaymentStatus3, map[string]interface{}{})
	if err != nil {
		uc.TX.Rollback()
		res = viewmodel.PaymentNotificationVm{
			Response:      "Payment Notification",
			TransactionID: transaction.ID,
			TrxID:         input.TrxID,
			MerchantID:    input.MerchantID,
			Merchant:      input.Merchant,
			BillNo:        input.BillNo,
			ResponseCode:  "00",
			ResponseDesc:  "Sukses",
			ResponseDate:  now,
		}

		return res, err, http.StatusOK
	}

	res = viewmodel.PaymentNotificationVm{
		Response:      "Payment Notification",
		TransactionID: transaction.ID,
		TrxID:         input.TrxID,
		MerchantID:    input.MerchantID,
		Merchant:      input.Merchant,
		BillNo:        input.BillNo,
		ResponseCode:  "00",
		ResponseDesc:  "Sukses",
		ResponseDate:  now,
	}
	uc.TX.Commit()

	return res, err, http.StatusOK
}
