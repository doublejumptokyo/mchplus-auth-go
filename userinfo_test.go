package mchplus_auth

import (
	"testing"
)

func TestUserinfo(t *testing.T) {
	is := initializeTest(t)
	if accessToken == "" {
		t.Skip()
	}

	u, err := GetUserInfo(accessToken)
	is.Nil(err)

	print(*u)
}
