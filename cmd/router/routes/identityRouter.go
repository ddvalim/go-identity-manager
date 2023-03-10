package routes

import (
	identityController "go-identity-manager/cmd/controller"
	"net/http"
)

var IdentityRouter = []Route{
	{
		Method: http.MethodPost,
		URI:    "/identity",
		Func:   identityController.Create,
	},
}
