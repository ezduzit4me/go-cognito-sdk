package clients

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoClient interface {
	SignUp(email string, password string) (error, string)
}

type awsCognitoClient struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientId   string
}

func NewCognitoClient(cognitoRegion string, cognitoAppClientID string) CognitoClient {

	conf := &aws.Config{Region: aws.String(cognitoRegion)}

	sess, err := session.NewSession(conf)
	client := cognito.New(sess)

	if err != nil {
		panic((err))
	}

	return &awsCognitoClient{
		cognitoClient: client,
		appClientId:   cognitoAppClientID,
	}
}

func (ctx *awsCognitoClient) SignUp(email string, password string) (error, string) {

	return nil, "Hello Golang"
}
