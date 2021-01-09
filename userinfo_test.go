package mchplus_auth

import (
	"testing"
)

func TestUserinfo(t *testing.T) {
	if accessToken == "" {
		t.Skip("access token is not found")
	}
	u, err := GetUserInfo(accessToken)
	if err != nil {
		t.Fatal(err)
	}
	pp(t, *u)
}
