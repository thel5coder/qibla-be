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
	"qibla-backend/helpers/messages"
	"qibla-backend/helpers/str"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
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
	var client http.Client
	fmt.Printf(string(bodyPost))


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
	transactionUc := TransactionUseCase{UcContract:uc.UcContract}
	_,err = transactionUc.CountBy("","id",invoiceID)
	if err != nil {
		return res,err
	}

	return res,err
}
