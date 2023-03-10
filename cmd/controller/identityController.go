package identityController

import (
	"go-identity-manager/internal/identity"
	"net/http"
)

type IdentityController struct {
	Service identity.IdentityService
}

func NewController(service identity.IdentityService) *IdentityController {
	return &IdentityController{
		Service: service,
	}
}

func Create(w http.ResponseWriter, r *http.Request) {}
