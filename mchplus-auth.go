package mchplus_auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lestrrat-go/jwx/jwa"
)

const alg = jwa.RS256

var (
	AuthAPI      = "https://auth.mch.plus/api"
	ClientID     = ""
	ClientSecret = ""
	RedirectURI  = ""
	Client       = &ClientInfo{}
)

type ClientInfo struct {
	ClientID    string `json:"client_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LogoURL     string `json:"logo_url"`
	HomeURL     string `json:"home_url"`
}

func ParseIDToken(idToken string, now int64) (p *Payload, err error) {
	p, err = ParseVerify(idToken)
	if err != nil {
		return
	}

	return p, CheckValidIDToken(p, now, Client.HomeURL)
}

func CheckValidIDToken(p *Payload, now int64, homeURL string) error {
	if p.IssuedAt().Unix() > now {
		return errors.New("id token error: issueAt")
	}

	if p.Expiration().Unix() < now {
		return errors.New("id token error: expire")
	}

	for _, a := range p.Audience() {
		if a == homeURL {
			return nil
		}
	}

	return errors.New("id token error: audience")
}

func Init(clientID, clientSecret, redirectURI string) (err error) {
	ClientID = clientID
	ClientSecret = clientSecret
	RedirectURI = redirectURI
	Client, err = GetClient()
	return err
}

func GetClient() (*ClientInfo, error) {
	b, err := get("/client?client_id="+ClientID, "")
	if err != nil {
		return nil, err
	}

	ret := new(ClientInfo)
	return ret, json.Unmarshal(b, ret)
}

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
