package ports

import "github.com/dianrahmaji/star-vault/internal/entity"

type IDNGateway interface {
	FetchLivestreams() ([]*entity.Livestream, error)
}

type YoutubeGateway interface {
	UploadVideo(livestream *entity.Livestream) (string, error)
	UpdateThumbnail(videoID string, thumbnailPath string) error
	CreatePlaylist(title string) (string, error)
	AddVideoToPlaylist(videoID, playlistID string) error
}
