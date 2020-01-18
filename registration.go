package mchplus_auth

import "encoding/json"

func RegisterRegion(address, alpha3 string) (err error) {
	in := map[string]string{
		"address":       address,
		"alpha3":        alpha3,
		"client_id":     ClientID,
		"client_secret": ClientSecret,
	}
	b, err := json.Marshal(in)
	if err != nil {
		return
	}
	_, err = post("/region/register", b)
	if err != nil {
		return
	}
	return
}

func RegisterPhone(address, phoneNumber string) (err error) {
	in := map[string]string{
		"address":      address,
		"phone_number": phoneNumber,
	}
	b, err := json.Marshal(in)
	if err != nil {
		return
	}
	_, err = post("/phone/register", b)
	if err != nil {
		return
	}
	return
}

func ConfirmPhone(address, signature, network string) (err error) {
	in := map[string]string{
		"address": address,
		"sig":     signature,
	}
	b, err := json.Marshal(in)
	if err != nil {
		return
	}
	_, err = post("/phone/confirm", b)
	if err != nil {
		return
	}

	return
}
