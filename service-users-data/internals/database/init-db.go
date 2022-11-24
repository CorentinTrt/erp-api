// Package database regroup all database's related stuff
package database

import "go.mongodb.org/mongo-driver/mongo"

const (
	dbURI = "mongodb://localhost:27017"
)

// MgCl is the mongodb client used by the service
var MgCl *Mongo

// Database is the main database of the service
var Database *mongo.Database

// UsersCol is the collection that regroup users's data
var UsersCol *mongo.Collection

func createDatabase(mg *Mongo) {
	Database = mg.Database("service-users-data")
}

func createCollections(mg *Mongo) {
	UsersCol = Database.Collection("users")
}

// InitDatabase will create a new MongoDB client, then connect that client and finally ping the server to check
// if everything's fine
func InitDatabase() error {
	mgCl, err := NewMongoClient(dbURI)
	if err != nil {
		return err
	}

	err = mgCl.ConnectClient()
	if err != nil {
		return err
	}

	err = mgCl.PingServer()
	if err != nil {
		return err
	}

	MgCl = mgCl

	createDatabase(MgCl)
	createCollections(MgCl)

	return nil
}
