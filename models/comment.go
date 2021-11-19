package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	ID        string               `json:"id,omitempty" bson:"_id,omitempty"`
	Message   string               `json:"message,omitempty" bson:"message"`
	Post      primitive.ObjectID   `json:"post,omitempty" bson:"post"`
	User      Author               `json:"user,omitempty" bson:"user"`
	CreatedAt time.Time            `json:"created_at" bson:"created_at"`
	Likes     []primitive.ObjectID `json:"likes,omitempty" bson:"likes"`
}
