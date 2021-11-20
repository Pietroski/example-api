package mockGenerator

// internal
//go:generate mockgen -package mockedSignUpFactory -destination ../../../internal/mocks/internal/factories/auth/signup/lambda/mockedSignUpFactory.go github.com/SimpleOpenBadge/authentication-api/internal/factories/auth/signup/lambda SignUpLambdaFactory
//go:generate mockgen -package mockedSignUpStore -destination ../../../internal/mocks/internal/services/datastore/dynamodb/auth/signup/mockedStore.go github.com/SimpleOpenBadge/authentication-api/internal/services/datastore/dynamodb/auth/signup Store

// external
//go:generate mockgen -package mockedFactory -destination ../../../pkg/mocks/pkg/models/data/aws/mockedFactory.go github.com/SimpleOpenBadge/authentication-api/pkg/models/data/aws LambdaFactory
