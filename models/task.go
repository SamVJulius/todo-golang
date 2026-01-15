package models

import "time"

type Task struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}
