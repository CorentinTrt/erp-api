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

// MongoBasic extends crud functionality and adds client functions
type MongoBasic interface {
	PingServer() error
	ConnectClient() error
}

// NewMongoClient is a constructor for a mongo client, it has to connect outside
func NewMongoClient(mgConStr string) (*Mongo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mgConStr))
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
