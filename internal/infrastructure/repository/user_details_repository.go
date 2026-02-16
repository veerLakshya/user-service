package repository

import (
	"github.com/veerlakshya/user-service/consts"
	"go.mongodb.org/mongo-driver/mongo"
)

type userDetailsRepository struct {
	BaseRepository
	collection *mongo.Collection
}

func createUserDetailsIndex(db *mongo.Database) {

}

func NewUserDetailsRepository(db *mongo.Database) *userDetailsRepository {
	createUserDetailsIndex(db)
	return &userDetailsRepository{
		BaseRepository: BaseRepository{
			db: db,
		},
		collection: db.Collection(consts.UserDetailsCollection),
	}
}
