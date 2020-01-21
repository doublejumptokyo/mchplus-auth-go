package mchplus_auth

import (
	"encoding/json"
)

type inviteCodeOutput struct {
	InviteCode   string `json:"invite_code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type addressOutput struct {
	Address      string `json:"address"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func GetAddressFromInviteCode(inviteCode string) (address string, err error) {
	in := map[string]string{
		"invite_code":   inviteCode,
		"client_id":     ClientID,
		"client_secret": ClientSecret,
	}
	b, err := json.Marshal(in)
	if err != nil {
		return
	}
	b, err = post("/referral/address", b)
	if err != nil {
		return
	}
	o := new(addressOutput)
	err = json.Unmarshal(b, o)
	if err != nil {
		return
	}
	address = o.Address
	return
}

func GetInviteCodeFromAddress(address string) (inviteCode string, err error) {
	in := map[string]string{
		"address":       address,
		"client_id":     ClientID,
		"client_secret": ClientSecret,
	}
	b, err := json.Marshal(in)
	if err != nil {
		return
	}
	b, err = post("/referral/code", b)
	if err != nil {
		return
	}
	o := new(inviteCodeOutput)
	err = json.Unmarshal(b, o)
	if err != nil {
		return
	}
	inviteCode = o.InviteCode
	return
}
