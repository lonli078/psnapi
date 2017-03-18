package psnapi

import (
	"bytes"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

)

const (
	sso_url  = "https://auth.api.sonyentertainmentnetwork.com/2.0/ssocookie"
	code_url = "https://auth.api.sonyentertainmentnetwork.com/2.0/oauth/authorize?"
	oath_url = "https://auth.api.sonyentertainmentnetwork.com/2.0/oauth/token"
)

type Api struct {
	Username   string
	Password   string
	npsso      string
	grand_code string
	oauth      string
	reflesh    string
	client     *http.Client
	reconnect  int
}

func (a *Api) Auth() {
	a.client = &http.Client{}
	a.get_sso()
	a.get_code()
	a.get_oauth()
}

func (a *Api) get_sso() {
	m := make(map[string]string)
	m["authentication_type"] = "password"
	m["username"] = a.Username
	m["password"] = a.Password
	m["client_id"] = "71a7beb8-f21a-47d9-a604-2e71bee24fe0"
	r, _ := http.NewRequest("POST", sso_url, bytes.NewBufferString(a.make_params(m)))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := a.client.Do(r)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	data := make(map[string]string)
	json.Unmarshal(body, &data)
	a.npsso = data["npsso"]
	fmt.Println(a.npsso)
}

func (a *Api) get_code() {
	m := make(map[string]string)
	m["state"] = "06d7AuZpOmJAwYYOWmVU63OMY"
	m["duid"] = "0000000d000400808F4B3AA3301B4945B2E3636E38C0DDFC"
	m["app_context"] = "inapp_ios"
	m["client_id"] = "b7cbf451-6bb6-4a5a-8913-71e61f462787"
	m["scope"] = "capone:report_submission,psn:sceapp,user:account.get,user:account.settings.privacy.get,user:account.settings.privacy.update,user:account.realName.get,user:account.realName.update,kamaji:get_account_hash,kamaji:ugc:distributor,oauth:manage_device_usercodes"
	m["response_type"] = "code"
	r, _ := http.NewRequest("Get", code_url+a.make_params(m), nil)
	r.AddCookie(&http.Cookie{Name: "npsso", Value: a.npsso})
	resp, _ := a.client.Do(r)
	defer resp.Body.Close()
	fmt.Println(resp.Header)
	a.grand_code = resp.Header["X-Np-Grant-Code"][0]
}

func (a *Api) get_oauth() {
	m := make(map[string]string)
	m["app_context"] = "inapp_ios"
	m["client_id"] = "b7cbf451-6bb6-4a5a-8913-71e61f462787"
	m["client_secret"] = "zsISsjmCx85zgCJg"
	m["code"] = a.grand_code
	m["duid"] = "0000000d000400808F4B3AA3301B4945B2E3636E38C0DDFC"
	m["grant_type"] = "authorization_code"
	m["scope"] = "capone:report_submission,psn:sceapp,user:account.get,user:account.settings.privacy.get,user:account.settings.privacy.update,user:account.realName.get,user:account.realName.update,kamaji:get_account_hash,kamaji:ugc:distributor,oauth:manage_device_usercodes"
	r, _ := http.NewRequest("POST", oath_url, bytes.NewBufferString(a.make_params(m)))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := a.client.Do(r)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	data := make(map[string]string)
	json.Unmarshal(body, &data)
	a.oauth = data["access_token"]
	a.reflesh = data["refresh_token"]
}

func (a *Api) make_params(params map[string]string) string {
	r := url.Values{}
	for k, v := range params {
		r.Add(k, v)
	}
	return r.Encode()
}

func (a *Api) make_get_request(url string, data *map[string]interface{}) error {
	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+a.oauth)
	resp, err := a.client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}

