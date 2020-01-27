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
	ok, address, err := GetAddressFromInviteCode(testInviteCode)
	is.Nil(err)
	if !ok {
		t.Errorf("ok is false")
	}
	if address != testAddress {
		t.Errorf("got %v\nwant %v", address, testAddress)
	}
}

func TestGetInviteCodeFromAddress(t *testing.T) {
	is := initializeTest(t)
	if testAddress == "" {
		t.Skip()
	}
	inviteCode, err := GetInviteCodeFromAddress(testAddress)
	is.Nil(err)
	if inviteCode != testInviteCode {
		t.Errorf("got %v\nwant %v", inviteCode, testInviteCode)
	}
}
