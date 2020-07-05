package fragment

import "time"

type DateTimeRange struct {
	StartedAt   time.Time `json:"startedAt"`
	EndedAt     time.Time `json:"endedAt"`
}