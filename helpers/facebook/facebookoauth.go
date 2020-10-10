package facebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OauthFacebookURLAPI ...
const OauthFacebookURLAPI = "https://graph.facebook.com/me?fields=id,name,email&access_token="

// GetFacebookProfile ...
func GetFacebookProfile(token string) (res map[string]interface{}, err error) {
	response, err := http.Get(OauthFacebookURLAPI + token)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	if response.StatusCode >= 400 {
		return res, errors.New("invalid_facebook_access_token")
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return res,err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return res, errors.New("error_read_body")
	}
	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, err
}
