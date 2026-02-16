package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepository struct {
	db *mongo.Database
}

type Database struct {
	DB                     *mongo.Database
	ClientCollection       *mongo.Collection
	UserDetailsCollection  *mongo.Collection
	RegisterUserCollection *mongo.Collection
	BrandKitCollection     *mongo.Collection
}

// Close closes the databse connection
func (db *Database) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return db.DB.Client().Disconnect(ctx)
}

func (r *BaseRepository) BeginTx(ctx context.Context) (mongo.SessionContext, error) {
	// start a new session
	session, err := r.db.Client().StartSession()
	if err != nil {
		return nil, err
	}

	// start a transaction
	err = session.StartTransaction()
	if err != nil {
		session.EndSession(ctx)
		return nil, err
	}

	// create a new context with session
	sessionCtx := mongo.NewSessionContext(ctx, session)
	return sessionCtx, nil
}
