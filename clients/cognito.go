package clients

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoClient interface {
	SignUp(email string, password string) (error, string)
	ConfirmSignUp(email string, code string) (error, string)
	// SignIn(email string, code string) (error, string)
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
	user := &cognito.SignUpInput{
		Username: aws.String(email),
		Password: aws.String(password),
		ClientId: aws.String(ctx.appClientId),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}
	result, err := ctx.cognitoClient.SignUp(user)
	if err != nil {
		return err, ""
	}
	return nil, result.String()
}

func (ctx *awsCognitoClient) ConfirmSignUp(email string, code string) (error, string) {
	confirmSignUpInput := &cognito.ConfirmSignUpInput{
		Username:         aws.String(email),
		ConfirmationCode: aws.String(code),
		ClientId:         aws.String(ctx.appClientId),
	}
	result, err := ctx.cognitoClient.ConfirmSignUp(confirmSignUpInput)

	if err != nil {

		return err, ""
	}

	return nil, result.String()

}

//func (ctx *awsCognitoClient) SignIn(email string, code string) (error, string) {
//	initiateAuthInput := &cognito.InitiateAuthInput()

//	result, err := ctx.cognitoClient.InitiateAuth()

//return nil, result.String()
// }
