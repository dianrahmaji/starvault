package entity

import "time"

type Streamer struct {
	StreamerID           int
	StreamerExternalUUID string
	Name                 *string
	IsLive               bool
	IsRecording          bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
