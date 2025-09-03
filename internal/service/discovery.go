package service

import (
	"fmt"

	"github.com/dianrahmaji/star-vault/internal/ports"
)

type DiscoveryService struct {
	IDNGateway ports.IDNGateway
}

func NewDiscoveryService(IDNGateway ports.IDNGateway) *DiscoveryService {
	return &DiscoveryService{IDNGateway: IDNGateway}
}

func (s *DiscoveryService) DiscoverStreamers() {
	livestreams, err := s.IDNGateway.FetchLivestreams()
	if err != nil {
		fmt.Println("error")
	}

	for _, ls := range livestreams {
		fmt.Printf("%+v\n", ls)
	}
}
