package usecase

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func (uc FaspayUseCase) PostData(input viewmodel.FaspayPostDataVm) (res map[string]interface{}, err error) {
	compose := os.Getenv("FASPAY_USER_ID") + `` + os.Getenv("FASPAY_PASSWORD") + `` + input.BillNo
	signature := uc.getSignature(compose)
	input.Signature = signature
	bodyPost, _ := json.Marshal(input)
	var client http.Client

	request, err := http.NewRequest("POST", fasPayBaseUrl+"/300011/10", bytes.NewBuffer(bodyPost))
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
	if err != nil {
		return res, err
	}

	return res, err
}
