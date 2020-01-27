package mchplus_auth

import (
	"encoding/json"
)

func SetClientSince(address, since string) (err error) {
	in := map[string]string{
		"address":       address,
		"since":         since,
		"client_id":     ClientID,
		"client_secret": ClientSecret,
	}
	b, err := json.Marshal(in)
	if err != nil {
		return
	}
	_, err = post("/client/since", b)
	return
}
