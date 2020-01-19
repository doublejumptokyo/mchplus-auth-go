package mchplus_auth

import (
	"testing"
)

func TestUserinfo(t *testing.T) {
	is := initializeTest(t)

	u, err := GetUserInfo(accessToken)
	is.Nil(err)

	print(u)
}
