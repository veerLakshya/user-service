package entity

import (
	"time"

	shared "github.com/veerlakshya/user-service/internal/domain/shared/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
}

type UserDetails struct {
	ID        primitive.ObjectID `bson:"id" json:"id"`
	Email     string             `bson:"email" json:"email"`
	Name      string             `bson:"name" json:"name"`
	AvatarURL string             `bson:"avatar_url,omitempty" json:"avatar_url,omitempty"`
	EditedBy  *shared.EditedBy   `bson:"edited_by,omitempty" json:"edited_by,omitempty"`

	CreatedAt time.Duration `bson:"created_at" json:"created_at"`
	UpdatedAt time.Duration `bson:"updated_at" json:"updated_at"`
}
