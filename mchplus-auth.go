package mchplus_auth

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/pkg/errors"
)

const alg = jwa.RS256

var (
	AuthAPI      = "https://auth.mch.plus/api"
	ClientID     = ""
	ClientSecret = ""
	RedirectURI  = ""
	cached       = map[string]*Metadata{}
)

func get(path string) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	resp, err := http.Get(AuthAPI + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Backend returns status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func post(path string) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	resp, err := http.Get(AuthAPI + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Backend returns status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Init(clientID, clientSecret, redirectURI string) (err error) {
	ClientID = clientID
	ClientSecret = clientSecret
	RedirectURI = redirectURI
	body, err := get("/metadata/x509")
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
