package mchplus_auth

import (
	"testing"
)

func TestGetAddressFromInviteCode(t *testing.T) {
	if testInviteCode == "" {
		t.Skip("test invite code is not found")
	}
	ok, address, err := GetAddressFromInviteCode(testInviteCode)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("ok is false")
	}
	if address != testAddress {
		t.Fatalf("got %v\nwant %v", address, testAddress)
	}
}

func TestGetInviteCodeFromAddress(t *testing.T) {

	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				"0xd868711BD9a2C6F1548F5f4737f71DA67d821090",
			},
			want:    "THA4",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, err := GetInviteCodeFromAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInviteCodeFromAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
			if code != tt.want {
				t.Errorf("GetInviteCodeFromAddress() code = %v, wantErr %v", code, tt.want)
			}
		})
	}
}
