// Package services regroup all services
// Services : entities that handle CRUD operation
package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	DB "service-users-data/internals/database"
	M "service-users-data/internals/database/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Users type is a slice of User
type Users = M.Users

// GetUsers function send to the client the list of all users
func GetUsers(resWriter http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cursor, err := DB.UsersCol.Find(ctx, bson.M{})
	if err != nil {
		http.Error(resWriter, "Database Error - Database can't be reach", http.StatusServiceUnavailable)
	}

	var u Users
	if err = cursor.All(ctx, &u); err != nil {
		fmt.Println(err)
		http.Error(resWriter, "Internal Server Error - Unable to process users data", http.StatusInternalServerError)
	}

	u.ToJSON(resWriter)
}
