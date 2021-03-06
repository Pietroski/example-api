package controllerResponses

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

// APIResponse processes and returns a lambda api response.
func APIResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		Headers: map[string]string{"Content-Type": "application/json"},
	}
	response.StatusCode = status

	stringBody, _ := json.Marshal(body)
	response.Body = string(stringBody)
	return &response, nil
}
