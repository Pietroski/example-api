package errorsModel

const (
	ErrorMethodNotAllowed = "method not allowed"

	ErrorFailedToFetchRecord = "failed to fetch record"

	ErrorCouldNotDynamoPutItem = "could not dynamo put item error"

	ErrorInvalidEmail         = "invalid email"
	ErrorInvalidPassword      = "invalid password"
	ErrorFailedToHashPassword = "failed to hash password"
	ErrorInvalidIssuerData    = "invalid issuer data"

	ErrorIssuerAlreadyExists = "issuer already exists"
	ErrorIssuerDoesNotExist  = "issuer does not exist"

	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"

	ErrorFailedToQueryUnescaped = "failed to query unescaped"
)
