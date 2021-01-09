package mchplus_auth

import (
	"encoding/json"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	if user == nil {
		t.Skip("user is not found")
	}
	t.Log("address: " + user.Address())

	msg, state, err := Authorize(user.Address())
	if err != nil {
		t.Error(err)
	}

	sig, err := user.PersonalSign(msg)
	if err != nil {
		t.Error(err)
	}
	t.Log("sig: " + sig)

	code, err := Login(sig, user.Address(), state, "mainnet")
	if err != nil {
		t.Error(err)
	}

	token, err := GetToken(code)
	if err != nil {
		t.Error(err)
	}
	accessToken = token.AccessToken
	t.Log("access token: " + accessToken)

	p, err := ParseIDToken(token.IDToken, time.Now().Unix())
	if err != nil {
		t.Error(err)
	}
	b, err := json.Marshal(p)
	if err != nil {
		t.Error(err)
	}
	t.Log("verified id_token: " + string(b))
}
