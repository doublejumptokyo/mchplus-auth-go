package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/doublejumptokyo/mchplus-auth-go"
)

var id = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjEiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOlsiaHR0cHM6Ly9nZ2cubWNoLnBsdXMiXSwiZXhwIjoxNTc0MTQ4NTc2LCJpYXQiOjE1NzQwNjIxNzYsImlzcyI6Ik1DSCsiLCJzdWIiOiIweGQ4Njg3MTFiZDlhMmM2ZjE1NDhmNWY0NzM3ZjcxZGE2N2Q4MjEwOTAifQ.rUr0O0DgViKh9xjoUIfjJeRtlECBWG86ZYc4ySWKMp8J9IEiaR997IuDds4CKwjGRISk57TM_4gkEracGqzyi-h5D6Pfc_pNM8ErK4zMyPmKdWGSLD_McF4LnMBZ0tmpUpGnuwE1KlP9nbKd7IEVkRcTftfjYHo00l72kkOdkee7xk_3xSpA5QJ0FkQo_O5XIC4t50LUlKCZsLkCS9Lmu2ooxXoQpI691hyPUNAZn5jWjp5P_I3FVjN9-_3t_7KgsbOJVWGtS_Fc4OE5vYpbmBFY_KFyrbne2nN5SzCtf14k-c_WD3evF1NxcieWtlHJuH0mFHTwx8H5r1XmJaUOVRY_EHE_ou9hU6Ka7k426rlDEcJgfkwRwvu18RrTobPhTs_N9NziShB2Tg-agIQIIZu5C8xz2E8uWPFeASGf2oDiWhVM8sdyoX5sNCfPDISuzquuKhEScK5LAteBGgc_ZDlNKUoF_VbPgXXtCWNfppS30pIgTwrh9hoGMI9VWW2Lc6uavUSI-9xcZCHIDR9QxJUe04yVma2IqiMJGLXmkOk2nLDrKngNznxEA9Kb-8OnQgQUqTZ8V2nBIsmiKonEKK7yz8ouHsyL0yprnPG0TR-dVWHjBmwNX_2L-yIoWVdD5FNyFULISHr6iJNjr-zvWawkmLNPTUle4DRS5D881nE"

func run(args []string) (err error) {
	k, err := mchplus_auth.ParseVerify(id)
	if err != nil {
		return
	}
	b, _ := json.Marshal(k)
	fmt.Printf("%v\n", string(b))
	return
}

func main() {
	flag.Parse()
	err := mchplus_auth.Init()
	if err != nil {
		panic(err)
	}
	if err := run(flag.Args()); err != nil {
		panic(err)

	}
}
