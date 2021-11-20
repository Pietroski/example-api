package factoryHandlers

import (
	"github.com/SimpleOpenBadge/authentication-api/pkg/controllers/responses"
	errorsModel "github.com/SimpleOpenBadge/authentication-api/pkg/models/domain/errors"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

// UnhandledMethod handles unwanted http methods by the default handler action.
func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return controllerResponses.APIResponse(
		http.StatusMethodNotAllowed, errorsModel.ErrorMethodNotAllowed,
	)
}
