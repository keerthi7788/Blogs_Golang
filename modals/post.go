package modals

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title_Post string             `json:"title,omitempty" bson:"title,omitempty"`
	Content    string             `json:"content,omitempty" bson:"content,omitempty"`
	AuthorID   primitive.ObjectID `json:"authorid,omitempty" bson:"authorid,omitempty"`
	CreatedAt  *time.Time         `json:"createdat,omitempty" bson:"createdat,createdat"`
}
