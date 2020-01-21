package mchplus_auth

import (
	"encoding/json"
)

type UserInfo struct {
	Address    string `json:"address"`
	InviteCode string `json:"invite_code"`
	PhoneHash  string `json:"phone_hash,omitempty"`
	Region     string `json:"region,omitempty"`
}

func GetUserInfo(accessToken string) (u *UserInfo, err error) {
	u = new(UserInfo)
	b, err := get("/userinfo", accessToken)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, u)
	return
}
