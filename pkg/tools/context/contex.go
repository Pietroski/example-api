package pietroski

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type (
	Context interface {
		ShouldBindWithJSON(body string, obj interface{}) error
	}
)

type (
	NetHttpContext struct {
		ctx     context.Context
		request http.Request
	}

	//LambdaRestContext struct {
	//	ctx     context.Context
	//	request events.APIGatewayProxyRequest
	//}

	LambdaWebSocketContext struct {
		ctx     context.Context
		request events.APIGatewayWebsocketProxyRequest
	}
)

//func (c *LambdaRestContext) ShouldBindWithJSON(body string, obj interface{}) error {
//	return binders.ShouldBindJSON(body, obj)
//}
