// Package database regroup all database's related stuff
package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo abstract the mongo client to use the functions
type Mongo struct {
	*mongo.Client
}

// MongoCRUD is an interface to abstract and test the packages
// type MongoCRUD interface {
// 	FindMany(database, collection string, selector bson.D, options options.FindOptions, output interface{}) error
// 	FindOne(database, collection string, selector bson.D, options options.FindOptions, output interface{}) error
// 	InsertOne(database, collection string, options options.InsertOneOptions, insert interface{}) error
// InsertMany(database, collection string, options options.InsertManyOptions, insert interface{}) error
// Upsert(database, collection string, selector bson.D, options options.UpdateOptions, update interface{}) error
// Update(database, collection string, selector bson.D, options options.UpdateOptions, update interface{}) error
// Delete(database, collection string, options options.DeleteOptions, selector bson.D) error
// }

// MongoBasic extends crud functionality and adds client functions
type MongoBasic interface {
	PingServer() error
	ConnectClient() error
	// MongoCRUD
}

// NewMongoClient is a constructor for a mongo client, it has to connect outside
func NewMongoClient(mongoConn string) (*Mongo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoConn))
	if err != nil {
		return nil, err
	}
	mg := Mongo{
		client,
	}
	return &mg, nil
}

// ConnectClient connects the client instance
func (m *Mongo) ConnectClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := m.Connect(ctx)

	return err
}

// PingServer pings the server and check if its found
func (m *Mongo) PingServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := m.Ping(ctx, readpref.Primary())

	return err
}
