package modals

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bosn:"_id,omitempty"`
	UserID    primitive.ObjectID
	Content   string
	CreatedAt primitive.DateTime
	PostID    primitive.ObjectID
}
