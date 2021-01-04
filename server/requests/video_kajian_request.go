package requests

type VideoKajianRequest struct {
	YoutubeVideoID string      `json:"youtube_video_id"`
	VideoContentID string      `json:"video_content_id"`
	Title          string      `json:"title"`
	ChannelName    string      `json:"channel_name"`
	Description    string      `json:"description"`
	EmbeddedPlayer string      `json:"embedded_player"`
	Thumbnails     interface{} `json:"thumbnails"`
	PublishedAt    string      `json:"published_at"`
}
