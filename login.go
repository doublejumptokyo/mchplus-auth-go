package mchplus_auth

import (
	"encoding/json"
	"net/url"

	"github.com/doublejumptokyo/mchplus-auth-go/utils"
	"github.com/pkg/errors"
)

type loginInput struct {
	Address     string `json:"address"`
	ClientID    string `json:"client_id"`
	Signature   string `json:"signature"`
	Network     string `json:"network"`
	RedirectURI string `json:"redirect_uri"`
	Lang        string `json:"lang"`
	State       string `json:"state"`
}

func Login(signature, address, state, network string) (code string, err error) {
	in := map[string]string{
		"address":      address,
		"client_id":    ClientID,
		"signature":    signature,
		"network":      network,
		"redirect_uri": RedirectURI,
		"state":        state,
	}
	inb, err := json.Marshal(in)
	if err != nil {
		return
	}
	b, err := post("/login", inb)
	if err != nil {
		return
	}

	res := map[string]string{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}

	if state != res["state"] {
		return "", errors.New("Invalid state")
	}
	return res["code"], nil

}

func Authorize(address string) (msg string, state string, err error) {

	state = utils.RandNumberString(6)
	q := make(url.Values)
	q.Set("response_type", "code")
	q.Set("client_id", ClientID)
	q.Set("state", state)
	q.Set("scope", "openid")
	q.Set("redirect_uri", RedirectURI)
	q.Set("address", address)
	b, err := get("/authorize?"+q.Encode(), "")
	if err != nil {
		return
	}

	res := map[string]string{}
	err = json.Unmarshal(b, &res)

	return res["message"], state, err
}
