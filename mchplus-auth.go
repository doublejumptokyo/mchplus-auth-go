package mchplus_auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/pkg/errors"
)

const MetadataAPI = "https://auth.mch.plus/api/metadata/x509"
const alg = jwa.RS256

var cached = map[string]*Metadata{}

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

func GetHeaders(idToken string) (*Headers, error) {
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

	h, err := GetHeaders(idToken)
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

func Init() (err error) {
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
