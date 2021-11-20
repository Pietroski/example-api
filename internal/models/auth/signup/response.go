package signupModel

import "time"

type (
	IssuerResponseModel struct {
		Email      string `json:"email"`
		IssuerName string `json:"issuerName"`

		IssuerVerified bool      `json:"issuerVerified"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}
)
