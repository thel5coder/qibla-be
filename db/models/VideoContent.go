package models

import "database/sql"

type VideoContent struct {
	ID        string         `db:"id"`
	Channel   string         `db:"channel"`
	ChannelID string         `db:"channel_id"`
	Links     string         `db:"links"`
	IsActive  bool           `db:"is_active"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
