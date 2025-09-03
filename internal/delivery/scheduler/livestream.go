package scheduler

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"

	"github.com/dianrahmaji/star-vault/config"
	"github.com/dianrahmaji/star-vault/internal/gateway"
	"github.com/dianrahmaji/star-vault/internal/repository/db/postgres"
	"github.com/dianrahmaji/star-vault/internal/service"
)

func Start(cfg *config.Config) error {
	ig := gateway.NewIDNGateway(cfg.IDNApiUrl)

	db, err := openDB(cfg.DSN)
	if err != nil {
		return err
	}

	sr := postgres.NewStreamerRepo(db)

	ds := service.NewDiscoveryService(ig, sr)

	ds.DiscoverStreamers()
	c := cron.New()

	c.AddFunc("@every 1m", func() {
		ds.DiscoverStreamers()
	})

	c.Start()

	select {}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
