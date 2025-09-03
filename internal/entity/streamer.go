package entity

import "time"

type Streamer struct {
	StreamerID           int
	StreamerPlatformUUID string
	Name                 string
	shouldRecord         bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
