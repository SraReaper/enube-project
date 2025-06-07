package allrouter

import (
	"SraReaper/enube-me-api/api/src/controllers"
	"net/http"
)

var userRoutes = []AllRoute{

	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		Authentication: false,
	},

	{
		URI:            "/users/{userId}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		Authentication: false,
	},
}
