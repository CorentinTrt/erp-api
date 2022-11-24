// Package api regroup all http related files
package api

import (
	"net/http"

	S "service-users-data/internals/services"
)

// UserH (handler) attach the needed services
type UserH struct {
}

// GetUsers function return a list of all users
func (userH *UserH) GetUsers(resWriter http.ResponseWriter, req *http.Request) {
	S.GetUsers(resWriter, req)
}

// CreateUser function insert a new user in the database, based on the body of the function
func (userH *UserH) CreateUser(resWriter http.ResponseWriter, req *http.Request) {
	S.CreateUser(resWriter, req)
}
