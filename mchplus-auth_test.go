package mchplus_auth

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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

var (
	testAddress    = os.Getenv("ADDRESS")
	testInviteCode = os.Getenv("INVITE_CODE")
	testSince      = os.Getenv("CLIENT_SINCE")
)

func TestGetClient(t *testing.T) {
	is := initializeTest(t)
	c, err := GetClient()
	is.Nil(err)
	print(*c)
}

func initializeTest(t *testing.T) is.I {
	is := is.New(t)
	var err error

	if os.Getenv("AUTH_API") != "" {
		AuthAPI = os.Getenv("AUTH_API")
	}

	if os.Getenv("PRIVATE_KEY") != "" {
		user, err = signer.NewSignerFromHex(os.Getenv("PRIVATE_KEY"))
		is.Nil(err)
	} else {
		user = nil
	}

	err = Init(clientID, clientSecret, redirectURI)
	is.Nil(err)
	return is
}

func print(in interface{}) {
	if reflect.TypeOf(in).Kind() == reflect.Struct {
		in, _ = json.Marshal(in)
		in = string(in.([]byte))
	}
	fmt.Println(in)
}
