# mchplus-auth-go


# Usage
## Validate ID Token on server

```golang
import (
	"fmt"
	"os"
	"time"
)

func init() {
	_ = mchplus_auth.Init(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("REDIRECT_URI"),
	)
}

func getToken(string code) (accessToken string, idToken string, err error) {
    token, err := mchplus_auth.GetToken(code)
	return token.AccessToken, token.IDToken, err
}

func authMiddleware() error {
	var idToken string
	now := time.Now().Unix()
	p, err := mchplus_auth.ParseIDToken(idToken, now)
	if err != nil {
		return err
	}

	fmt.Printf("Hello address: ", p.Subject())
	return nil
}

func userInfo(accessToken string) (phoneHash string, error) {
    u, err := mchplus_auth.GetUserInfo(accessToken)
	return u.PhoneHash, err
}
```


# Testing

## Basic
```sh
$ export CLIENT_ID=your_client_id
$ export CLIENT_SECRET=your_client_secret
$ export REDIRECT_URI=your_redirect_uri
$ export PRIVATE_KEY=0x111

go test
```

## Phone register

```sh
$ go test -run TestRegisterPhone YOUR_PHONE_NUMBER
$ go test -run TestConfirmPhone CODE
```

