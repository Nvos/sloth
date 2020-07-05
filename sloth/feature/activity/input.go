package activity

import "time"

type ActivityCreateInput struct {
	StartedAt   time.Time `json:"startedAt"`
	EndedAt     time.Time `json:"endedAt"`
	Description string    `json:"description"`
}