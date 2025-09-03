package service

import (
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

	for _, ls := range livestreams {
		streamer := entity.Streamer{
			StreamerPlatformUUID: ls.Creator.UUID,
			Name:                 ls.Creator.Name,
		}

		err = s.StreamerRepo.Upsert(&streamer)
		if err != nil {
			return nil
		}
	}

	return nil
}
