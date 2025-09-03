package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/dianrahmaji/star-vault/internal/entity"
)

type StreamerRepo struct {
	DB *sql.DB
}

func NewStreamerRepo(db *sql.DB) *StreamerRepo {
	return &StreamerRepo{DB: db}
}

func (r *StreamerRepo) Upsert(streamer *entity.Streamer) error {
	query := `
		INSERT INTO streamers (platform_uuid, name)
		VALUES ($1, $2)
		ON CONFLICT (platform_uuid)
		DO UPDATE SET
			name = EXCLUDED.name,
			updated_at = now()
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, streamer.StreamerPlatformUUID, streamer.Name)
	if err != nil {
		return err
	}

	return nil
}
