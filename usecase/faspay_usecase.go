package usecase

type FaspayUseCase struct {
	*UcContract
}

func (uc FaspayUseCase) GetLisPaymentMethods() (res string, err error) {
	//compose := os.Getenv("FASPAY_USER_ID") + `` + os.Getenv("FASPAY_PASSWORD")
	//var md5 = md5.New()
	//var sha1 = sha1.New()
	//
	//md5.Write([]byte(compose))
	//md5EncryptedStr := md5.Sum(nil)
	//md5Str := fmt.Sprintf("%x",md5EncryptedStr)
	//
	//sha1.Write([]byte(md5Str))
	//sha1EncryptedStr := sha1.Sum(nil)
	//sh1Str := fmt.Sprintf("%x",sha1EncryptedStr)

	//var client http.Client
	//
	//sha1Encrypted := sha1.Sum(md5.Sum([]byte(compose)))
	//sha1EncryptedString := fmt.Sprintf("%x", sha1Encrypted)
	//var bodyPost = []byte(`{"merchant_id":"` + os.Getenv("FASPAY_MERCHANT_ID") + `","signature":"` + sha1EncryptedString + `"`)
	//
	//request, err := http.NewRequest("POST", fasPayBaseUrl+"/users", bytes.NewBuffer(bodyPost))
	//if err != nil {
	//	return res, err
	//}
	//
	//response, err := client.Do(request)
	//if err != nil {
	//	return res, err
	//}
	//defer response.Body.Close()
	//
	//body, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))

	return res, err
}