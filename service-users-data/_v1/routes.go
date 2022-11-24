// Package api regroup all http related files
package api

import (
	"github.com/go-chi/chi"
)

// UserRoutes function attach each route to the right handler
func UserRoutes() *chi.Mux {
	r := chi.NewRouter()

	userH := &UserH{}

	r.Get("/", userH.GetUsers)
	r.Post("/", userH.CreateUser)

	return r
}
