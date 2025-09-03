package ports

import "github.com/dianrahmaji/star-vault/internal/entity"

type CreatorDTO struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type LivestreamDTO struct {
	Title       string     `json:"title"`
	ImageUrl    string     `json:"image_url"`
	PlaybackUrl string     `json:"playback_url"`
	Status      string     `json:"status"`
	Creator     CreatorDTO `json:"creator"`
}

type IDNGateway interface {
	FetchLivestreams() ([]LivestreamDTO, error)
}

type YoutubeGateway interface {
	UploadVideo(livestream *entity.Livestream) (string, error)
	UpdateThumbnail(videoID string, thumbnailPath string) error
	CreatePlaylist(title string) (string, error)
	AddVideoToPlaylist(videoID, playlistID string) error
}
