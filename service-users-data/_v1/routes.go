// Package api regroup all http related files
package api

import (
	"github.com/go-chi/chi"
)

// UserRoutes function attach each route to the right handler
func UserRoutes() *chi.Mux {
	r := chi.NewRouter()

	userH := &UserH{}

	// swagger:route GET /users User listUsers
	// Return a list of all Users
	//
	// responses:
	//	200: usersResponse
	//  500: errorResponse
	//  503: errorResponse
	r.Get("/", userH.GetUsers)
	r.Post("/", userH.CreateUser)

	return r
}
