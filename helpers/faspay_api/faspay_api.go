package faspay_api

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

// Credential ...
type Credential struct {
	BaseURL             string
	UserID              string
	Password            string
	MerchantID          string
	DisbursementBaseURL string
	AppKey              string
	AppSecret           string
}

func getSignature(compose string) (res string) {
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

// func (cred Credential) Disbursement(input requests.FaspayPostRequest) (res map[string]interface{}, err error) {
// 	compose := os.Getenv("FASPAY_USER_ID") + `` + os.Getenv("FASPAY_PASSWORD") + `` + input.InvoiceNumber
// 	signature := uc.getSignature(compose)
// 	var client http.Client
// 	var items []viewmodel.ItemFaspayPostDataVm

// 	for _, item := range input.Item {
// 		items = append(items, viewmodel.ItemFaspayPostDataVm{
// 			Product:     item.Product,
// 			Amount:      item.Amount,
// 			Qty:         item.Qty,
// 			PaymentPlan: item.PaymentPlan,
// 			Tenor:       item.Tenor,
// 			MerchantID:  os.Getenv("FASPAY_MERCHANT_ID"),
// 		})
// 	}
// 	body := viewmodel.FaspayPostDataVm{
// 		Request:        input.RequestTransaction,
// 		MerchantID:     os.Getenv("FASPAY_MERCHANT_ID"),
// 		Merchant:       os.Getenv("FASPAY_MERCHANT"),
// 		BillNo:         input.InvoiceNumber,
// 		BillDate:       input.TransactionDate,
// 		BillExpired:    input.DueDate,
// 		BillDesc:       input.TransactionDesc,
// 		BillCurrency:   defaultFaspayCurrency,
// 		BillTotal:      input.Total,
// 		PaymentChannel: input.PaymentChannel,
// 		PayType:        defaultFaspayPayType,
// 		CustNo:         input.UserID,
// 		CustName:       input.CustomerName,
// 		Msisdn:         str.StringToInt(input.CustomerPhoneNumber),
// 		Email:          input.CustomerEmail,
// 		Terminal:       defaultFaspayTerminal,
// 		Signature:      signature,
// 		Item:           items,
// 	}
// 	bodyPost, _ := json.Marshal(body)

// 	request, err := http.NewRequest("POST", os.Getenv("FASPAY_BASE_URL")+"/300011/10", bytes.NewBuffer(bodyPost))
// 	if err != nil {
// 		return res, err
// 	}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		return res, err
// 	}
// 	defer response.Body.Close()

// 	responseBody, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return res, err
// 	}
// 	err = json.Unmarshal(responseBody, &res)
// 	if err != nil {
// 		return res, err
// 	}
// 	fmt.Println(res)
// 	if res["response_code"] != "00" {
// 		return res, errors.New(messages.PaymentFailed)
// 	}

// 	return res, err
// }
