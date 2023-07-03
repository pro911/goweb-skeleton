package mao

import "go.mongodb.org/mongo-driver/bson"

type TaskTestEvents struct {
	TaskID        string   `bson:"task_id" json:"task_id"`
	EventID       string   `bson:"event_id" json:"event_id"`
	ParentEventID string   `bson:"parent_event_id" json:"parent_event_id"`
	ProjectID     string   `bson:"project_id" json:"project_id"`
	Type          string   `bson:"type" json:"type"`
	Enabled       int64    `bson:"enabled" json:"enabled"`
	Sort          int64    `bson:"sort" json:"sort"`
	Data          bson.Raw `bson:"data" json:"data"`
}
