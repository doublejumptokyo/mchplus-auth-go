package mchplus_auth

import (
	"fmt"
	"os"
	"testing"

	"github.com/cheekybits/is"
	"github.com/doublejumptokyo/mchplus-auth-go/utils/signer"
)

var (
	clientID     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	redirectURI  = os.Getenv("REDIRECT_URI")
)

var (
	user        = &signer.Signer{}
	accessToken = os.Getenv("ACCESS_TOKEN")
)

func TestUserinfo(t *testing.T) {
	is := initializeTest(t)
	is.Nil(nil)
	print(accessToken)
}

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

func initializeTest(t *testing.T) is.I {
	is := is.New(t)
	var err error

	if os.Getenv("AUTH_API") != "" {
		AuthAPI = os.Getenv("AUTH_API")
	}

	user, err = signer.NewSignerFromHex(os.Getenv("PRIVATE_KEY"))
	is.Nil(err)

	err = Init(clientID, clientSecret, redirectURI)
	is.Nil(err)
	return is
}

func print(in interface{}) {
	fmt.Println(in)
}
