package awsModel

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type (
	// AWSPack is the simplest most common aws service pack used for this project factories.
	AWSPack struct {
		AwsConfig    *aws.Config
		DynamoClient *dynamodb.Client
	}

	// LambdaFactory is the simplest most common aws interface used for this project factories.
	LambdaFactory interface {
		Handler(
			ctx context.Context,
			request events.APIGatewayProxyRequest,
		) (*events.APIGatewayProxyResponse, error)
		Start()
	}
)
