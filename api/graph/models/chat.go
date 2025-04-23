package models

import "time"

type Chat struct {
	ID          ObjectID   `json:"id" bson:"_id"`
	Title       string     `json:"title"`
	Model       string     `json:"model"`
	Messages    []*Message `json:"messages"`
	LastChanged time.Time  `json:"last_changed"`
}
