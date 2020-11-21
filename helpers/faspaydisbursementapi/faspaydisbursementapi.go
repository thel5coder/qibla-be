package faspaydisbursementapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"qibla-backend/helpers/interfacepkg"
	"strings"
	"time"
)

var (
	getTokenURL = "/account/api/tokens"
	registerURL = "/account/api/register"
)

// Credential ...
type Credential struct {
	BaseURL      string
	Key          string
	Secret       string
	AppKey       string
	AppSecret    string
	ClientKey    string
	ClientSecret string
	IV           string
	SourceVA     string
}

// getAuthorization ...
func (cred *Credential) getAuthorization() string {
	plaintext := cred.AppKey + ":" + cred.AppSecret
	key := GeneratePassword(cred.Secret)
	firstIv := GenerateIv(cred.IV)

	// Initialize new crypter struct. Errors are ignored.
	crypter, _ := NewCrypter(key, firstIv)

	// Lets encode plaintext using the same key and iv.
	// This will produce the very same result: "RanFyUZSP9u/HLZjyI5zXQ=="
	encoded, _ := crypter.Encrypt([]byte(plaintext))
	return base64.StdEncoding.EncodeToString(encoded)
}

// getSignature ...
func (cred *Credential) getSignature(method, path, timestamp, body string) string {
	encodedStr := ""
	if body != "" {
		// SHA256 body
		h := sha256.New()
		h.Write([]byte(body))
		bodySha := h.Sum(nil)

		// Hex and lowercase body
		src := []byte(bodySha)
		encodedStr = hex.EncodeToString(src)
		encodedStr = strings.ToLower(encodedStr)
	}

	token := base64.StdEncoding.EncodeToString([]byte(cred.ClientKey + ":" + cred.ClientSecret))
	stringToSign := method + ":" + path + ":" + timestamp + ":" + token + ":" + encodedStr
	fmt.Println(stringToSign)
	secret := []byte(cred.Secret)
	message := []byte(stringToSign)

	hash := hmac.New(sha256.New, secret)
	hash.Write(message)

	// to lowercase hexits
	return hex.EncodeToString(hash.Sum(nil))
}

// GenerateToken ...
func (cred *Credential) GenerateToken() (res map[string]interface{}, err error) {
	fullURL := cred.BaseURL + getTokenURL

	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := cred.getSignature(method, getTokenURL, timestamp, "")
	fmt.Println(signature)
	authorization := cred.getAuthorization()
	fmt.Println(authorization)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest(method, fullURL, nil)
	r.Header.Add("faspay-key", cred.Key)
	r.Header.Add("faspay-timestamp", timestamp)
	r.Header.Add("faspay-signature", signature)
	r.Header.Add("faspay-authorization", authorization)

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

// Register ...
func (cred *Credential) Register(data map[string]interface{}) (res map[string]interface{}, err error) {
	fullURL := cred.BaseURL + registerURL

	// Make payload
	payload := map[string]interface{}{
		"virtual_account":          cred.SourceVA,
		"beneficiary_account":      data["beneficiary_account"].(string),
		"beneficiary_account_name": data["beneficiary_account_name"].(string),
		"beneficiary_va_name":      data["beneficiary_va_name"].(string),
		"beneficiary_bank_code":    data["beneficiary_bank_code"].(string),
		"beneficiary_bank_branch":  data["beneficiary_bank_branch"].(string),
		"beneficiary_region_code":  data["beneficiary_region_code"].(string),
		"beneficiary_country_code": data["beneficiary_country_code"].(string),
		"beneficiary_purpose_code": data["beneficiary_purpose_code"].(string),
	}

	method := "POST"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := cred.getSignature(method, registerURL, timestamp, interfacepkg.Marshall(payload))
	fmt.Println(signature)
	authorization := cred.getAuthorization()
	fmt.Println(authorization)

	// token := base64.StdEncoding.EncodeToString([]byte(cred.ClientKey + cred.ClientSecret))
	// stringToSign := method + ":" + registerURL + ":" + timestamp + ":" + token + ":" + interfacepkg.Marshall(payload) + ":" + interfacepkg.Marshall(payload)
	// tokenAPI, err := cred.GenerateToken("POST", registerURL, timestamp, token, interfacepkg.Marshall(payload), stringToSign)
	// fmt.Println(tokenAPI, err)

	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest(method, fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("faspay-key", cred.Key)
	r.Header.Add("faspay-timestamp", timestamp)
	r.Header.Add("faspay-signature", signature)
	r.Header.Add("faspay-authorization", authorization)
	r.Header.Add("Content-Type", "application/json")

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
