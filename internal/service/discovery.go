package service

import (
	"slices"

	"github.com/dianrahmaji/star-vault/internal/entity"
	"github.com/dianrahmaji/star-vault/internal/ports"
)

type DiscoveryService struct {
	IDNGateway   ports.IDNGateway
	StreamerRepo ports.StreamerRepository
}

func NewDiscoveryService(IDNGateway ports.IDNGateway, streamerRepo ports.StreamerRepository) *DiscoveryService {
	return &DiscoveryService{IDNGateway: IDNGateway, StreamerRepo: streamerRepo}
}

func (s *DiscoveryService) DiscoverStreamers() error {
	livestreams, err := s.IDNGateway.FetchLivestreams()
	if err != nil {
		return err
	}

	var streamers = make([]*entity.Streamer, 0)

	for _, ls := range livestreams {
		streamer := entity.Streamer{
			StreamerPlatformUUID: ls.Creator.UUID,
			Name:                 ls.Creator.Name,
		}

		i := slices.IndexFunc(streamers, func(s *entity.Streamer) bool {
			if s == nil {
				return false
			}

			return s.StreamerPlatformUUID == streamer.StreamerPlatformUUID
		})

		if i == -1 {
			streamers = append(streamers, &streamer)
		}
	}

	err = s.StreamerRepo.UpsertMany(streamers)
	if err != nil {
		return err
	}

	return nil
}
