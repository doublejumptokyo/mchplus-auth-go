package mchplus_auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

var (
	cachedMetadata = map[string]*Metadata{}
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

func GetToken(code string) (*Token, error) {
	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Set("code", code)
	values.Set("redirect_uri", RedirectURI)
	values.Set("client_id", ClientID)
	values.Set("client_secret", ClientSecret)
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
	ikid, _ := h.Get("kid")
	kid := ikid.(string)

	m := new(Metadata)
	ok := false
	m, ok = cachedMetadata[kid]
	if !ok {
		err = fetchMetadata()
		if err != nil {
			return nil, err
		}
		m, ok = cachedMetadata[kid]
		if !ok {
			return nil, errors.New("kid not found")
		}
	}

	v, err := jws.Verify([]byte(idToken), alg, m.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to verify jws signature: %w", err)
	}

	p := new(Payload)
	if err := json.Unmarshal(v, p); err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return p, nil
}

func fetchMetadata() (err error) {
	body, err := get("/metadata/x509", "")
	if err != nil {
		return
	}
	res := map[string]string{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	for k, v := range res {
		block, _ := pem.Decode([]byte(v))
		if block == nil {
			return errors.New("invalid public key data")
		}
		var err error
		c := new(Metadata)
		c.PublicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return err
		}
		cachedMetadata[k] = c
	}
	return
}
