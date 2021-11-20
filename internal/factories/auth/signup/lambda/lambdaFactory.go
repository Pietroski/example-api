package signUpFactory

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"net/http"

	signUpController "github.com/SimpleOpenBadge/authentication-api/internal/controllers/auth/signup/lambda"
	"github.com/SimpleOpenBadge/authentication-api/internal/services/datastore/dynamodb/auth/signup"
	"github.com/SimpleOpenBadge/authentication-api/pkg/factories/handlers"
	"github.com/SimpleOpenBadge/authentication-api/pkg/models/data/aws"
)

type SignUpLambdaFactory interface {
	awsModel.LambdaFactory
}

// SignUpFactory factory's struct model.
type SignUpFactory struct {
	awsPack          *awsModel.AWSPack
	store            signup.Store
	signUpController *signUpController.SignUpController
}

// NewSignUpFactory instantiates a new SignUp factory.
func NewSignUpFactory(
	awsConfig *aws.Config,
	store signup.Store,
) SignUpLambdaFactory {
	signUpFactory := &SignUpFactory{
		awsPack: &awsModel.AWSPack{
			AwsConfig:    awsConfig,
			DynamoClient: nil,
		},
		signUpController: signUpController.NewSignUpController(store),
	}

	return signUpFactory
}

// Handler handles http methods accordingly...
func (suf *SignUpFactory) Handler(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case http.MethodPost:
		return suf.signUpController.CreateIssuer(ctx, request)
	case http.MethodPatch:
		return suf.signUpController.UpdateIssuer(ctx, request)
	case http.MethodDelete:
		return suf.signUpController.DeleteIssuer(ctx, request)
	case http.MethodGet:
		// TODO: switch between GetIssuer and ListIssuers...
		return suf.signUpController.GetIssuer(ctx, request)

	default:
		return factoryHandlers.UnhandledMethod()
	}
}

// Start initiates lambda server.
func (suf *SignUpFactory) Start() {
	lambda.Start(suf.Handler)
}
