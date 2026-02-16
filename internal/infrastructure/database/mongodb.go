package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDb(ctx context.Context, uri string, dbName string) (*MongoDB, error) {
	// set client options
	clientOptions := options.Client().ApplyURI(uri)

	// connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// ping mongoDB server to verify connection
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	// get database
	db := client.Database(dbName)
	return &MongoDB{
		client: client,
		db:     db,
	}, nil
}

// Close disconnects from MongoDB
func (m *MongoDB) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

// Collection returns a MongoDB collection
func (m *MongoDB) Collection(name string) *mongo.Collection {
	return m.db.Collection(name)
}

// Client returns the MongoDB client
func (m *MongoDB) Client() *mongo.Client {
	return m.client
}

// Database returns the MongoDB database
func (m *MongoDB) Database() *mongo.Database {
	return m.db
}
