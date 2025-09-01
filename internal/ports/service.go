package ports

import "github.com/dianrahmaji/star-vault/internal/entity"

type DiscoveryService interface {
	DiscoverStreamers() error

	GetStreamers() ([]*entity.Streamer, error)
	GetStreamerById(id int) (*entity.Streamer, error)
	GetStreamerbyPlatformUUID(uuid string) (*entity.Streamer, error)

	SaveLivestream(livestream *entity.Livestream) error
	UpdateLivestream(livestream *entity.Livestream) error
	GetLivestreams() ([]*entity.Livestream, error)
	GetLivestreamByID() (*entity.Livestream, error)
}

type RecordingService interface {
	StartRecording(livestream *entity.Livestream) error
}

type UploadService interface {
	Upload(livestream *entity.Livestream) error
}
