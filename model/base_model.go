package model

import (
	"example/todo/util"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt *time.Time          `json:"created_at" bson:"createdAt,omitempty"`
	UpdatedAt *time.Time          `json:"updated_at" bson:"updatedAt,omitempty"`
}

func (base *Base) Init() {
	timestamp := util.GetCurrentUTCTime()
	objID := primitive.NewObjectID()

	base.ID = &objID
	base.CreatedAt = timestamp
	base.UpdatedAt = timestamp
}
