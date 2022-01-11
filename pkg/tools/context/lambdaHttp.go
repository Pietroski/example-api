package pietroski

import (
	"context"
	"github.com/SimpleOpenBadge/authentication-api/pkg/tools/binders"
	"github.com/SimpleOpenBadge/authentication-api/pkg/tools/validators"
	"github.com/aws/aws-lambda-go/events"
)

type LambdaRestContext struct {
	ctx     context.Context
	request events.APIGatewayProxyRequest
}

func (ctx *LambdaRestContext) ShouldBindWithJSON(obj interface{}) (err error) {
	if err = binders.ShouldBindJSON(ctx.request.Body, obj); err != nil {
		return
	}

	return validators.NewValidator(obj).Validate()
}
