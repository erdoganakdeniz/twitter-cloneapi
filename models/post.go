package models

import (
	"github.com/erdoganakdeniz/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PostInput struct {
	Title       string `json:"title,omitempty" bson:"title" valid:"length(3|30)"`
	Description string `json:"description,omitempty" bson:"description" valid:"length(3|300)"`
}
type Post struct {
	ID          string               `json:"id,omitempty" bson:"_id"`
	Title       string               `json:"title,omitempty" bson:"title"`
	Description string               `json:"description,omitempty" bson:"description"`
	CreatedAt   time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at" bson:"updated_at"`
	Author      Author               `json:"author,omitempty" bson:"author"`
	Comments    []primitive.ObjectID `json:"comments,omitempty" bson:"comments"`
	Likes       []primitive.ObjectID `json:"likes,omitempty" bson:"likes"`
}
type PostWithComment struct {
	ID          string               `json:"id,omitempty" bson:"_id"`
	Title       string               `json:"title,omitempty" bson:"title"`
	Description string               `json:"description,omitempty" bson:"description"`
	CreatedAt   time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at" bson:"updated_at"`
	Author      Author               `json:"author,omitempty" bson:"author"`
	Comments    []Comment `json:"comments,omitempty" bson:"comments"`
	Likes       []primitive.ObjectID `json:"likes,omitempty" bson:"likes"`
}

func (i PostInput) Validate() error {
	return utils.Validator(i)
}