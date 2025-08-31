package entity

import "time"

type LivestreamStatus string

const (
	LivestreamStatusLive      = "live"
	LivestreamStatusScheduled = "scheduled"
)

type Livestream struct {
	LivestreamID int
	StreamerID   int
	Title        string
	ThumbnailUrl string
	PlaybackUrl  string
	Status       LivestreamStatus
	StartedAt    *time.Time
	EndedAt      *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
