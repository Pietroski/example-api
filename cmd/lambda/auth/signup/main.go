package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/SimpleOpenBadge/authentication-api/internal/factories/auth/signup/lambda"
	"github.com/SimpleOpenBadge/authentication-api/internal/services/datastore/dynamodb/auth/signup"
)

func main() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-west-2"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	dynamodbClient := dynamodb.NewFromConfig(cfg)
	signUpStore := signup.NewStore(dynamodbClient)
	authSignUpServer := signUpFactory.NewSignUpFactory(&cfg, signUpStore)
	authSignUpServer.Start()
}
