// Package services regroup all services
// Services : entities that handle CRUD operation
package services

import (
	"fmt"
	"net/http"

	M "service-users-data/internals/database/models"
)

// CreateUser function insert a new user in the database
func CreateUser(resWriter http.ResponseWriter, req *http.Request) {
	nwUser := &M.User{}

	err := nwUser.FromJSON(req.Body)
	if err != nil {
		http.Error(resWriter, "Internal Server Error - Unable to unmarshal data", http.StatusInternalServerError)
	}

	fmt.Printf("input: %#v", nwUser)
}
