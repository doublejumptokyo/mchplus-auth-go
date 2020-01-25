package mchplus_auth

import (
	"flag"
	"testing"
)

func TestRegisterRegion(t *testing.T) {
	is := initializeTest(t)
	if user == nil {
		t.Skip()
	}
	args := flag.Args()
	if len(args) == 0 {
		t.Skip()
	}
	region := args[0]
	print(user.Address())
	err := RegisterRegion(user.Address(), region)
	is.Nil(err)
}

func TestRegisterPhone(t *testing.T) {
	is := initializeTest(t)
	if user == nil {
		t.Skip()
	}

	print(user.Address())
	args := flag.Args()
	if len(args) == 0 {
		t.Skip()
	}
	number := args[0]
	print(number)
	err := RegisterPhone(user.Address(), number)
	is.Nil(err)
}

func TestConfirmPhone(t *testing.T) {
	is := initializeTest(t)
	if user == nil {
		t.Skip()
	}

	print(user.Address())
	args := flag.Args()
	if len(args) == 0 {
		t.Skip()
	}
	code := args[0]
	sig, err := user.PersonalSign("Code:" + code)
	is.Nil(err)
	err = ConfirmPhone(user.Address(), sig, "mainnet")
	is.Nil(err)
}
