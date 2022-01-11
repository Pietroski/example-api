package signup

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"time"

	errorsModel "github.com/SimpleOpenBadge/authentication-api/pkg/models/domain/errors"
)

const (
	issuerTableName = "IssuerTable"
)

func (ds *DynamoDBStore) GetIssuer(ctx context.Context, issuerEmail string) (*Issuer, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{
				Value: issuerEmail,
			},
		},
		TableName: aws.String(issuerTableName),
	}

	result, err := ds.dynamodbClient.GetItem(ctx, input)
	if err != nil {
		return nil, errors.New(errorsModel.ErrorFailedToFetchRecord)
	}

	item := new(Issuer)
	if err = attributevalue.UnmarshalMap(result.Item, item); err != nil {
		return nil, errors.New(errorsModel.ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

func (ds *DynamoDBStore) ListIssuers(ctx context.Context) ([]*Issuer, error) {
	// TODO: implement pagination

	input := &dynamodb.ScanInput{
		TableName: aws.String(issuerTableName),
	}

	result, err := ds.dynamodbClient.Scan(ctx, input)
	if err != nil {
		return nil, errors.New(errorsModel.ErrorFailedToFetchRecord)
	}

	items := new([]*Issuer)
	if err = attributevalue.UnmarshalListOfMaps(result.Items, items); err != nil {
		return nil, errors.New(errorsModel.ErrorFailedToUnmarshalRecord)
	}

	return *items, nil
}

func (ds *DynamoDBStore) StoreIssuer(ctx context.Context, issuer *_Issuer_) (*Issuer, error) {
	mav, err := attributevalue.MarshalMap(issuer)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      mav,
		TableName: aws.String(issuerTableName),
	}

	result, err := ds.dynamodbClient.PutItem(ctx, input)
	if err != nil {
		return nil, err
	}

	item := new(Issuer)
	if err = attributevalue.UnmarshalMap(result.Attributes, item); err != nil {
		return nil, errors.New(errorsModel.ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

func (ds *DynamoDBStore) DeleteIssuer(ctx context.Context, issuerEmail string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{
				Value: issuerEmail,
			},
		},
		TableName: aws.String(issuerTableName),
	}

	_, err := ds.dynamodbClient.DeleteItem(ctx, input)
	if err != nil {
		return errors.New(errorsModel.ErrorCouldNotDeleteItem)
	}

	return nil
}

func (ds *DynamoDBStore) CreateIssuer(ctx context.Context, issuer *Issuer) (*Issuer, error) {
	existingIssuer, _ := ds.GetIssuer(ctx, issuer.Email)
	if existingIssuer != nil && len(existingIssuer.Email) != 0 {
		return nil, errors.New(errorsModel.ErrorIssuerAlreadyExists)
	}

	i := _Issuer_{
		Email:          issuer.Email,
		IssuerName:     issuer.IssuerName,
		Password:       issuer.Password,
		IssuerVerified: false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return ds.StoreIssuer(ctx, &i)
}

func (ds *DynamoDBStore) UpdateIssuer(ctx context.Context, issuer *Issuer) (*Issuer, error) {
	existingIssuer, _ := ds.GetIssuer(ctx, issuer.Email)
	if existingIssuer == nil && len(existingIssuer.Email) == 0 {
		return nil, errors.New(errorsModel.ErrorIssuerDoesNotExist)
	}

	// TODO: verify this role
	i := _Issuer_{
		Email:          issuer.Email,
		IssuerName:     issuer.IssuerName,
		Password:       issuer.Password,
		IssuerVerified: false,
		CreatedAt:      issuer.CreatedAt,
		UpdatedAt:      time.Now(),
	}

	return ds.StoreIssuer(ctx, &i)
}
