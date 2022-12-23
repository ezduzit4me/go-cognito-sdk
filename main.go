package main

import (
	"fmt"

	"github.com/ezduzit4me/go-cognito-sdk/clients"
)

func main() {

	cognitoClient := clients.NewCognitoClient("", "")
	err, result := cognitoClient.SignUp("geoffreyWarrener@gmail.com", "test")

	if err != nil {
		panic(err)
	}
	fmt.Printf("result:" + fmt.Sprintf("%s", result))
}
