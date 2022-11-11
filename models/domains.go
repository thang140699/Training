package models

import "time"

type SetTime struct {
	Time   time.Time `json:"Time" bson:"Time"`
	Domain string    `json:"Domain" bson:"Domain"`
}
