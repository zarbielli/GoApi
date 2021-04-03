package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		HTTPMethod:  http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		HTTPMethod:  http.MethodGet,
		Function:    controllers.GetAllUsers,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		HTTPMethod:  http.MethodGet,
		Function:    controllers.GetOneUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		HTTPMethod:  http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		HTTPMethod:  http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
