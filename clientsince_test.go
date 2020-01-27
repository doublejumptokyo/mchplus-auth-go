package mchplus_auth

import (
	"testing"
)

func TestSetClientSince(t *testing.T) {
	is := initializeTest(t)
	if testInviteCode == "" {
		t.Skip()
	}
	err := SetClientSince(testAddress, testSince)
	is.Nil(err)
}
