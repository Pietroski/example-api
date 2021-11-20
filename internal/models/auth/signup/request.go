package signupModel

type (
	IssuerRequestCreationModel struct { // all required
		Email      string `json:"email"`
		IssuerName string `json:"issuerName"`
		Password   string `json:"password"`
	}

	IssuerRequestUpdateModel struct {
		Email          string `json:"email"` // required
		IssuerName     string `json:"issuerName,omitempty"`
		Password       string `json:"password,omitempty"`
		IssuerVerified bool   `json:"issuerVerified,omitempty"`
	}
)
