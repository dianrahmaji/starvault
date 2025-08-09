package main

import "time"

type User struct {
	ID          string `gorm:"primaryKey"`
	Name        *string
	IsLive      bool
	IsRecording bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OAuthToken struct {
	ID           uint `gorm:"primaryKey"`
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type LiveStreamStatus string

const (
	StatusLive      LiveStreamStatus = "live"
	StatusScheduled LiveStreamStatus = "scheduled"
)

type LiveStreamCreator struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type LiveStreamMetadata struct {
	Folder string `json:"folder"`
}

type LiveStream struct {
	Title       string             `json:"title"`
	ImageUrl    string             `json:"image_url"`
	PlaybackUrl string             `json:"playback_url"`
	Status      LiveStreamStatus   `json:"status"`
	Creator     LiveStreamCreator  `json:"creator"`
	Metadata    LiveStreamMetadata `json:"metadata"`
}
