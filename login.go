package mchplus_auth

import (
	"encoding/json"
	"net/url"

	"github.com/doublejumptokyo/mchplus-auth-go/utils"
)

func Authorize(address string) (msg string, state string, err error) {

	state = utils.RandNumberString(6)
	q := make(url.Values)
	q.Set("response_type", "code")
	q.Set("client_id", ClientID)
	q.Set("state", state)
	q.Set("scope", "openid")
	q.Set("redirect_uri", RedirectURI)
	b, err := get("/authorize?" + q.Encode())
	if err != nil {
		return
	}

	res := map[string]string{}
	err = json.Unmarshal(b, &res)

	return res["message"], state, err
}
