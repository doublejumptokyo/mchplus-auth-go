package mchplus_auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/pkg/errors"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}

type Headers struct {
	jws.StandardHeaders
}

type Payload struct {
	jwt.Token
}

type Metadata struct {
	PublicKey *rsa.PublicKey
}

func GetToken(clientID, clientSecret, redirectURI, code string) (*Token, error) {
	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Set("code", code)
	values.Set("redirect_uri", redirectURI)
	values.Set("client_id", clientID)
	values.Set("client_secret", clientSecret)
	tokenAPI := AuthAPI + "/token"
	resp, err := http.PostForm(tokenAPI, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Backend returns status %d", resp.StatusCode)
	}

	t := new(Token)
	return t, json.Unmarshal(body, t)
}

func ParseHeaders(idToken string) (*Headers, error) {
	raw := strings.Split(idToken, ".")[0]

	b, err := base64.RawStdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}

	h := new(Headers)
	err = json.Unmarshal(b, h)
	return h, err
}

func ParseVerify(idToken string) (*Payload, error) {

	h, err := ParseHeaders(idToken)
	if err != nil {
		return nil, err
	}
	kid, ok := h.Get("kid")
	if !ok {
		return nil, errors.New("kid not found")
	}

	v, err := jws.Verify([]byte(idToken), alg, cached[kid.(string)].PublicKey)
	if err != nil {
		return nil, errors.Wrap(err, `failed to verify jws signature`)
	}

	p := new(Payload)
	if err := json.Unmarshal(v, p); err != nil {
		return nil, errors.Wrap(err, `failed to parse token`)
	}

	return p, nil
}
