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
	_, err = post("/permission/client/since", b)
	return
}

func GetUserinfoPermissioned(address string) (u *UserInfo, err error) {
	in := map[string]string{
		"address":       address,
		"client_id":     ClientID,
		"client_secret": ClientSecret,
	}
	b, err := json.Marshal(in)
	if err != nil {
		return
	}
	b, err = post("/permission/userinfo", b)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, u)
	return
}
