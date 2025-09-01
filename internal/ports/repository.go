package ports

import "github.com/dianrahmaji/star-vault/internal/entity"

type StreamerRepository interface {
	Save(streamer *entity.Streamer) error
	Update(streamer *entity.Streamer) error
	FindAll() ([]*entity.Streamer, error)
	FindByID(id int) (*entity.Streamer, error)
	FindByPlatformUUID(uuid string) (*entity.Streamer, error)
}

type LivestreamRepository interface {
	Save(livestream *entity.Livestream) error
	Update(livestream *entity.Livestream) error
	FindAll() ([]*entity.Livestream, error)
	FindByID(id int) (*entity.Livestream, error)
}
