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
)

type FaspayUseCase struct {
	*UcContract
}

func (uc FaspayUseCase) GetLisPaymentMethods() (res map[string]interface{}, err error) {
	fmt.Print("ini")
	compose := os.Getenv("FASPAY_USER_ID") + `` + os.Getenv("FASPAY_PASSWORD")
	var md5 = md5.New()
	var sha1 = sha1.New()

	md5.Write([]byte(compose))
	md5EncryptedStr := md5.Sum(nil)
	md5Str := fmt.Sprintf("%x",md5EncryptedStr)

	sha1.Write([]byte(md5Str))
	sha1EncryptedStr := sha1.Sum(nil)
	sh1Str := fmt.Sprintf("%x",sha1EncryptedStr)

	var client http.Client

	//sha1Encrypted := sha1.Sum(md5.Sum([]byte(compose)))
	//sha1EncryptedString := fmt.Sprintf("%x", sha1Encrypted)
	var bodyPost = []byte(`{"merchant_id":"` + os.Getenv("FASPAY_MERCHANT_ID") + `","signature":"` + sh1Str + `"}`)

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
		return res,err
	}
	err = json.Unmarshal(body,&res)


	return res, err
}
