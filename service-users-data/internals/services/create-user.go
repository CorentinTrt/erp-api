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
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser function insert a new user in the database
func CreateUser(resWriter http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	nwUser := &M.User{}

	err := nwUser.FromJSON(req.Body)
	if err != nil {
		http.Error(resWriter, "Internal Server Error - Unable to unmarshal data", http.StatusInternalServerError)
		return
	}

	err = nwUser.Validate()
	if err != nil {
		http.Error(resWriter, fmt.Sprintf("Request Error - Wrong input: %s", err), http.StatusBadRequest)
		return
	}

	selector := bson.D{bson.E{Key: "email", Value: nwUser.Email}}
	result := DB.UsersCol.FindOne(ctx, selector)
	err = result.Decode(nwUser)

	if err == mongo.ErrNoDocuments {
		insertUser(ctx, nwUser)

		resWriter.Header().Set("Content-Type", "application/json")

		nwUser.ToJSON(resWriter)
		return
	} else if err != nil {
		http.Error(resWriter, fmt.Sprintf("Database Error - can't be reach: %s", err), http.StatusConflict)

		ctx.Done()
		return
	}

	http.Error(resWriter, fmt.Sprintf("Request Error - Duplicate user"), http.StatusConflict)

}

func insertUser(ctx context.Context, nwUser *M.User) {
	_, err := DB.UsersCol.InsertOne(ctx, nwUser)

	if err != nil {
		ctx.Done()
	}
}
