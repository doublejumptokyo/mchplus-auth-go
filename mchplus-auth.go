package mchplus_auth

import (
	"bytes"
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

func get(path string, authorization string) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	client := new(http.Client)
	req, err := http.NewRequest("GET", AuthAPI+path, nil)
	if authorization != "" {
		req.Header.Add("Authorization", "Bearer "+authorization)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Backend returns status %d msg: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func post(path string, body []byte) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	resp, err := http.Post(AuthAPI+path, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Backend returns status %d msg: %s", resp.StatusCode, string(ret))
	}

	return ret, nil
}

func Init(clientID, clientSecret, redirectURI string) (err error) {
	ClientID = clientID
	ClientSecret = clientSecret
	RedirectURI = redirectURI
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
		cached[k] = c
	}

	return nil
}
