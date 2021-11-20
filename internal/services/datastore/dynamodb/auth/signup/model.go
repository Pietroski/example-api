package signup

import (
	"context"
	"time"
)

type (
	Issuer struct {
		Email          string           `json:"email"`
		IssuerName     string           `json:"issuerName"`
		Password       string           `json:"password"`
		IssuerVerified bool             `json:"issuerVerified"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}

	_Issuer struct {
		Email          string    `json:"email"`
		IssuerName     string    `json:"issuerName"`
		Password       string    `json:"password"`
		IssuerVerified bool      `json:"issuerVerified"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}

	_Issuer_ struct {
		Email          string    `json:"email"`
		IssuerName     string    `json:"issuerName"`
		Password       string    `json:"password"`
		IssuerVerified bool      `json:"issuerVerified"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}

	Querier interface {
		GetIssuer(ctx context.Context, issuerEmail string) (*Issuer, error)
		ListIssuers(ctx context.Context) ([]*Issuer, error)
		DeleteIssuer(ctx context.Context, issuerEmail string) error

		StoreIssuer(ctx context.Context, issuer *_Issuer_) (*Issuer, error)
		CreateIssuer(ctx context.Context, issuer *Issuer) (*Issuer, error)
		UpdateIssuer(ctx context.Context, issuer *Issuer) (*Issuer, error)
		//ValidateIssuer(ctx context.Context, issuer *Issuer) (*Issuer, error)
	}
)
