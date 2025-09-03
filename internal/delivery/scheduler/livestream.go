package scheduler

import (
	"github.com/robfig/cron/v3"

	"github.com/dianrahmaji/star-vault/internal/gateway"
	"github.com/dianrahmaji/star-vault/internal/service"
)

func Start(IDNApiUrl string) {
	ig := gateway.NewIDNGateway(IDNApiUrl)
	ds := service.NewDiscoveryService(ig)

	ds.DiscoverStreamers()
	c := cron.New()

	c.AddFunc("@every 1m", func() {
		ds.DiscoverStreamers()
	})

	c.Start()

	select {}
}
