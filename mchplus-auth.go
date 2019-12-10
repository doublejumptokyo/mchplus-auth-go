package mchplus_auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/pkg/errors"
)

const alg = jwa.RS256

var (
	MetadataAPI  = "https://auth.mch.plus/api/metadata/x509"
	TokenAPI     = "https://auth.mch.plus/api/token"
	cached       = map[string]*Metadata{}
	clientID     = ""
	clientSecret = ""
	redirectURI  = ""
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

func Get() (map[string]*Metadata, error) {
	return cached, nil
}

func GetToken(code string) (*Token, error) {
	values := new(url.Values)
	values.Set("grant_type", "authorization_code")
	values.Set("code", code)
	values.Set("redirect_uri", redirectURI)
	values.Set("client_id", clientID)
	values.Set("client_secret", clientSecret)

	resp, err := http.Post(TokenAPI, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
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

func Init(id, secret, redirectURI string) (err error) {
	resp, err := http.Get(MetadataAPI)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
		cached[k] = c
	}

	return nil
}
