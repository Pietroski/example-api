package signup

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Store interface {
	Querier
}

type DynamoDBStore struct {
	dynamodbClient *dynamodb.Client
}

func NewStore(dynamodbClient *dynamodb.Client) Store {
	return &DynamoDBStore{
		dynamodbClient: dynamodbClient,
	}
}
