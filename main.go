package main

import (
	"fmt"

	"github.com/ezduzit4me/go-cognito-sdk/clients"
)

func main() {

	cognitoClient := clients.NewCognitoClient("ap-southeast-2", "26c2pgelskg6cq61htpck65ii3")
	//err, result, initiateAuthOutput := cognitoClient.SignUp("iran.rourke@dropsin.net", "5252/66Lam-4005")
	err, result, initiateAuthOutput := cognitoClient.SignIn("iran.rourke@dropsin.net", "5252/66Lam-4005")

	if err != nil {
		fmt.Print("Error" + fmt.Sprint("%s", err))
		panic(err)
	}
	fmt.Printf("Result:" + fmt.Sprintf("%s \n %s", result, *initiateAuthOutput.AuthenticationResult.IdToken))
}
