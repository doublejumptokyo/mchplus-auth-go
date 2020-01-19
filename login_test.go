package mchplus_auth

import (
	"testing"
)

func TestLogin(t *testing.T) {
	is := initializeTest(t)
	print(user.Address())
	msg, state, err := Authorize(user.Address())
	is.Nil(err)
	// print(msg)

	sig, err := user.PersonalSign(msg)
	is.Nil(err)
	// print(sig)

	code, err := Login(sig, user.Address(), state, "mainnet")
	is.Nil(err)
	// print(code)

	token, err := GetToken(code)
	is.Nil(err)
	accessToken = token.AccessToken
	print(accessToken)
}
