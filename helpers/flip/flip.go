package flip

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"qibla-backend/usecase/viewmodel"
	"strconv"
	"strings"
)

// Credential ...
type Credential struct {
	BaseURL         string
	SecretKey       string
	ValidationToken string
}

var (
	getBankURL      = "/general/banks"
	disbursementURL = "/disbursement"
)

// GetBank ...
func (cred *Credential) GetBank() (res []viewmodel.BankVM, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.SecretKey+":"))
	fullURL := cred.BaseURL + getBankURL

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("GET", fullURL, nil)
	r.Header.Add("Authorization", auth)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		fmt.Println(err)
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// Disbursement ...
func (cred *Credential) Disbursement(id, accountNumber, bankCode string, amount float64, remark, recipientCity string) (res viewmodel.DisbursementVM, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.SecretKey+":"))
	fullURL := cred.BaseURL + disbursementURL

	data := url.Values{}
	data.Set("account_number", accountNumber)
	data.Set("bank_code", bankCode)
	data.Set("amount", strconv.Itoa(int(amount)))
	data.Set("remark", remark)
	data.Set("recipient_city", recipientCity)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, strings.NewReader(data.Encode()))
	r.Header.Add("Authorization", auth)
	r.Header.Add("idempotency-key", id)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}
