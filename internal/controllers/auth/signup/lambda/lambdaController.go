package signUpController

import (
	"context"
	"errors"
	"fmt"
	signupModel "github.com/SimpleOpenBadge/authentication-api/internal/models/auth/signup"
	"github.com/SimpleOpenBadge/authentication-api/internal/services/datastore/dynamodb/auth/signup"
	controllerResponses "github.com/SimpleOpenBadge/authentication-api/pkg/controllers/responses"
	errorsModel "github.com/SimpleOpenBadge/authentication-api/pkg/models/domain/errors"
	patchesModel "github.com/SimpleOpenBadge/authentication-api/pkg/models/domain/patches"
	"github.com/SimpleOpenBadge/authentication-api/pkg/tools/authenticators/passwords"
	"github.com/SimpleOpenBadge/authentication-api/pkg/tools/binders"
	"github.com/SimpleOpenBadge/authentication-api/pkg/tools/validators"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"net/url"
)

// TODO: tag validators
// TODO: error wrapper

const (
	//
)

type SignUpController struct {
	store signup.Store
}

func NewSignUpController(store signup.Store) *SignUpController {
	return &SignUpController{
		store: store,
	}
}

func (suc *SignUpController) GetIssuer(
	ctx context.Context,
	r events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	// TODO: implement controllers
	// TODO: should contain Bearer accessToken in the headers??
	return &events.APIGatewayProxyResponse{}, nil
}

func (suc *SignUpController) ListIssuers(
	ctx context.Context,
	r events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	// TODO: implement controllers
	// TODO: should contain Bearer accessToken in the headers
	// TODO should be admin
	return &events.APIGatewayProxyResponse{}, nil
}

func (suc *SignUpController) DeleteIssuer(
	ctx context.Context,
	r events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	// TODO: implement controllers
	// TODO: should contain email in the payload
	// TODO: should contain Bearer accessToken in the headers
	return &events.APIGatewayProxyResponse{}, nil
}

func (suc *SignUpController) CreateIssuer(
	ctx context.Context,
	r events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	var irm signupModel.IssuerRequestCreationModel

	{ // Binders and Validators context.
		if err := binders.ShouldBindJSON(r.Body, irm); err != nil {
			return nil, err
		}
		if ok := validators.IsEmailValid(irm.Email); !ok {
			return nil, errors.New(errorsModel.ErrorInvalidEmail)
		}
		if ok := validators.IsPasswordValid(irm.Password); !ok {
			return nil, errors.New(errorsModel.ErrorInvalidPassword)
		}
	}

	hashedPassword, err := passwords.HashPassword(irm.Password)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", errorsModel.ErrorFailedToHashPassword, err)
	}

	// TODO: generate uuid?

	i := signup.Issuer{
		Email:          irm.Email,
		IssuerName:     irm.IssuerName,
		Password:       hashedPassword,
	}

	createdIssuer, err := suc.store.CreateIssuer(ctx, &i)
	if err != nil {
		return nil, err
	}

	return controllerResponses.APIResponse(http.StatusCreated, createdIssuer)
}

func (suc *SignUpController) UpdateIssuer(
	ctx context.Context,
	r events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	// TODO: implement controllers
	// TODO: should contain Bearer accessToken in the headers

	// TODO: verify authorisationToken
	// TODO: retrieve issuer
	// TODO: issuer must be verified to be updated.

	var irum signupModel.IssuerRequestUpdateModel
	{ // Binders and Validators context.
		if err := binders.ShouldBindJSON(r.Body, irum); err != nil {
			return nil, err
		}
	}

	raw, ok := r.QueryStringParameters[patchesModel.PatchItem]
	if ok {
		value, err := url.QueryUnescape(raw)
		if err != nil {
			return nil, errors.New(errorsModel.ErrorFailedToQueryUnescaped)
		}

		switch value {
		case patchesModel.PatchEmail:
			return suc.updateIssuerEmail(ctx, irum)
		case patchesModel.PatchIssuerName:
			return suc.updateIssuerName(ctx, irum)
		case patchesModel.PatchPassword:
			return suc.updateIssuerPassword(ctx, irum)
		case patchesModel.PatchIssuerVerification:
			return suc.updateIssuerVerification(ctx, irum)
		}
	}

	return nil, errors.New(patchesModel.PatchNotAllowedError)
}

func (suc *SignUpController) updateIssuerEmail(
	ctx context.Context,
	payload signupModel.IssuerRequestUpdateModel,
) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{}, nil
}

func (suc *SignUpController) updateIssuerName(
	ctx context.Context,
	payload signupModel.IssuerRequestUpdateModel,
) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{}, nil
}

func (suc *SignUpController) updateIssuerPassword(
	ctx context.Context,
	payload signupModel.IssuerRequestUpdateModel,
) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{}, nil
}

func (suc *SignUpController) updateIssuerVerification(
	ctx context.Context,
	payload signupModel.IssuerRequestUpdateModel,
) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{}, nil
}
