package mchplus_auth

import (
	"fmt"
	"os"
	"testing"

	"github.com/cheekybits/is"
)

var (
	clientID     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	redirectURI  = os.Getenv("REDIRECT_URI")
	address      = os.Getenv("ADDRESS")
)

var (
	accessToken = ""
)

func TestLogin(t *testing.T) {
	is := initializeTest(t)
	msg, _, err := Authorize(address)
	is.Nil(err)

	print(msg)
}

func initializeTest(t *testing.T) is.I {
	is := is.New(t)
	var err error

	if os.Getenv("AUTH_API") != "" {
		AuthAPI = os.Getenv("AUTH_API")
	}

	err = Init(clientID, clientSecret, redirectURI)
	is.Nil(err)
	return is
}

func print(in interface{}) {
	fmt.Println(in)
}
