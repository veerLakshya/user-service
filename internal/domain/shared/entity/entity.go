package entity

type EditedBy struct {
	ID    string `bson:"id" json:"id"`
	Email string `bson:"email" json:"email"`
}
