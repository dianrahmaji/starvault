package port

import "github.com/dianrahmaji/star-vault/internal/entity"

type StreamerRepository interface {
	UpsertMany(streamers []*entity.Streamer) error
}

type LivestreamRepository interface {
	Save(livestream *entity.Livestream) error
	Update(livestream *entity.Livestream) error
	FindAll() ([]*entity.Livestream, error)
	FindByID(id int) (*entity.Livestream, error)
}
