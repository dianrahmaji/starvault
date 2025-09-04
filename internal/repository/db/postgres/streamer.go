package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/dianrahmaji/star-vault/internal/entity"
)

type StreamerRepo struct {
	DB *sql.DB
}

func NewStreamerRepo(db *sql.DB) *StreamerRepo {
	return &StreamerRepo{DB: db}
}

func (r *StreamerRepo) UpsertMany(streamers []*entity.Streamer) error {
	if len(streamers) == 0 {
		return nil
	}

	var (
		values []string
		args   []any
	)

	for i, s := range streamers {
		if s == nil {
			continue
		}

		idx := i*2 + 1
		values = append(values, fmt.Sprintf("($%d, $%d)", idx, idx+1))
		args = append(args, s.StreamerPlatformUUID, s.Name)
	}

	query := fmt.Sprintf(`
		INSERT INTO streamers (platform_uuid, name)
		VALUES %s
		ON CONFLICT (platform_uuid)
		DO UPDATE SET
			name = EXCLUDED.name,
			updated_at = now()
		WHERE streamers.name IS DISTINCT FROM EXCLUDED.name;
	`, strings.Join(values, ","))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
