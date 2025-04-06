package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt *time.Time          `json:"created_at" bson:"createdAt,omitempty"`
	UpdatedAt *time.Time          `json:"updated_at" bson:"updatedAt,omitempty"`
}

type Todo struct {
	Base  `bson:",inline"`
	Title string `json:"title" bson:"title"`
	Done  bool   `json:"done" bson:"done"`
}

func (base *Base) Init() {
	t := time.Now().UTC()
	objID := primitive.NewObjectID()

	base.ID = &objID
	base.CreatedAt = &t
	base.UpdatedAt = &t
}

func (dest *Todo) Copy(src Todo) {
	dest.Title = src.Title
	dest.Done = src.Done
}
