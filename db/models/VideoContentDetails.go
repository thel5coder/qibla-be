package models

import (
	"database/sql"
	"time"
)

type VideoContentDetails struct {
	ID             string       `db:"id"`
	YoutubeVideoID string       `db:"youtube_video_id"`
	VideoContentID string       `db:"video_content_id"`
	Title          string       `db:"title"`
	ChannelTitle   string       `db:"channel_title"`
	EmbeddedUrl    string       `db:"embedded_url"`
	Thumbnails     string       `db:"thumbnails"`
	Description    string       `db:"description"`
	PublishedAt    string       `db:"published_at"`
	CreatedAt      time.Time    `db:"created_at"`
	UpdatedAt      time.Time    `db:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
}
