package mchplus_auth

import (
	"encoding/json"
)

type UserInfo struct {
	Address     string `json:"address"`
	PhoneHash   string `json:"phone_hash"`
	Region      string `json:"region"`
	FirstClient string `json:"first_client"`
	Since       int64  `json:"since"`
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
