package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SlackWebhooks struct {
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
}

type Organization struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	Identifier    string             `json:"identifier" bson:"identifier"`
	SlackWebhooks []SlackWebhooks    `json:"slackWebhooks" bson:"slackWebhooks"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"upadtedAt" bson:"updatedAt"`
}

func (o *Organization) Validate() bool {
	if o.Name == "" {
		return false
	}

	return true
}