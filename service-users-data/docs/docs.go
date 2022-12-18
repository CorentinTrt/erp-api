// Package classification Users' Data API
//
// Documentation for Users' Data API
//
//	Schemes: http
//	BasePath: /v1
//	Version: 0.1.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package classification

import (
	M "service-users-data/internals/database/models"
)

// A list of all Users
// swagger:response usersResponse
type productsResponseWrapper struct {
	// All current Users
	// in: body
	Body []M.User
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body M.GenericError
}
