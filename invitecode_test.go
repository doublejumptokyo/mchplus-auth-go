package mchplus_auth

import (
	"os"
	"testing"
)

var (
	testAddress    = os.Getenv("ADDRESS")
	testInviteCode = os.Getenv("INVITE_CODE")
)

func TestGetAddressFromInviteCode(t *testing.T) {
	is := initializeTest(t)
	if testInviteCode == "" {
		t.Skip()
	}
	address, err := GetAddressFromInviteCode(testInviteCode)
	is.Nil(err)
	print(address)
}

func TestGetInviteCodeFromAddress(t *testing.T) {
	is := initializeTest(t)
	if testAddress == "" {
		t.Skip()
	}
	inviteCode, err := GetInviteCodeFromAddress(testAddress)
	is.Nil(err)
	print(inviteCode)
}
