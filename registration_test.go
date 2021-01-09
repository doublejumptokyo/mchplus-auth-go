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

func TestRegisterBirthday(t *testing.T) {
	type args struct {
		address  string
		birthday string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				"0xd868711BD9a2C6F1548F5f4737f71DA67d821090",
				"1989-04-20",
			},
			wantErr: false,
		},
		{
			name: "wrong data",
			args: args{
				"0xd868711BD9a2C6F1548F5f4737f71DA67d821090",
				"2020-04-21",
			},
			wantErr: true,
		},
		{
			name: "parse error",
			args: args{
				"0xd868711BD9a2C6F1548F5f4737f71DA67d821090",
				"20200420",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterBirthday(tt.args.address, tt.args.birthday); (err != nil) != tt.wantErr {
				t.Errorf("RegisterBirthday() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
