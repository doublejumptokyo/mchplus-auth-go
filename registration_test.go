package mchplus_auth

import (
	"flag"
	"testing"
)

func TestRegisterRegion(t *testing.T) {
	is := initializeTest(t)
	print(user.Address())
	err := RegisterRegion(user.Address(), "JPN")
	is.Nil(err)
}

func TestRegisterPhone(t *testing.T) {
	is := initializeTest(t)
	print(user.Address())
	number := flag.Args()[0]
	print(number)
	err := RegisterPhone(user.Address(), number)
	is.Nil(err)
}

func TestConfirmPhone(t *testing.T) {
	is := initializeTest(t)
	print(user.Address())
	code := flag.Args()[0]
	sig, err := user.PersonalSign("Code:" + code)
	is.Nil(err)
	err = ConfirmPhone(user.Address(), sig, "mainnet")
	is.Nil(err)
}
