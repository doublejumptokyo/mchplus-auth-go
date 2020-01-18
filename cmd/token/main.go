package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/doublejumptokyo/mchplus-auth-go"
)

func run(args []string) (err error) {
	clientID := args[0]
	secret := args[1]
	redirectURI := args[2]
	accessToken := args[3]
	t, err := mchplus_auth.GetToken(clientID, secret, redirectURI, accessToken)
	if err != nil {
		return
	}
	b, _ := json.Marshal(t)
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
