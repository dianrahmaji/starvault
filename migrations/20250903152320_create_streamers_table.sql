-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS streamers (
	id bigserial PRIMARY KEY,
	platform_uuid TEXT NOT NULL UNIQUE,
	name TEXT NOT NULL,
	should_record BOOLEAN DEFAULT false,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF NOT EXISTS streamers;
-- +goose StatementEnd
