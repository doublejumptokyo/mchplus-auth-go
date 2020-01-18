package mchplus_auth

import (
	"encoding/json"
	"net/url"
)

func Authorize(address string) (msg string, state int64, err error) {

	q := make(url.Values)
	q.Set("response_type", "code")
	q.Set("client_id", ClientID)
	q.Set("state", "11111")
	q.Set("scope", "openid")
	q.Set("redirect_uri", RedirectURI)
	b, err := get("/authorize?" + q.Encode())
	if err != nil {
		return
	}

	res := map[string]string{}
	err = json.Unmarshal(b, &res)

	return res["message"], 11111, err
}
