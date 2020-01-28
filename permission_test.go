package mchplus_auth

import (
	"testing"
)

func TestSetClientSince(t *testing.T) {
	is := initializeTest(t)
	if testAddress == "" || testSince == "" {
		t.Skip()
	}
	err := SetClientSince(testAddress, testSince)
	is.Nil(err)
}

func TestGetUserinfoPermissioned(t *testing.T) {
	is := initializeTest(t)
	if testAddress == "" {
		t.Skip()
	}
	u, err := GetUserinfoPermissioned(testAddress)
	is.Nil(err)

	print(*u)
}
