package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt *time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt *time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
}

func (base *Base) Init() {
	t := time.Now().UTC()
	objID := primitive.NewObjectID()

	base.ID = &objID
	base.CreatedAt = &t
	base.UpdatedAt = &t
}

type LogEntry struct {
	Base        `bson:",inline"`
	Level       string `json:"level" bson:"level"`
	Message     string `json:"message" bson:"message"`
	ServiceName string `json:"serviceName" bson:"serviceName"`
	RequestID   string `json:"request_id,omitempty" bson:"request_id,omitempty"`
	Error       string `json:"error,omitempty" bson:"error,omitempty"`
}
