package viewmodel

type VideoDetailVm struct {
	ID             string                 `json:"id"`
	Title          string                 `json:"title"`
	ChannelName    string                 `json:"channel_name"`
	Description    string                 `json:"description"`
	EmbeddedPlayer string                 `json:"embedded_player"`
	Thumbnails     map[string]interface{} `json:"thumbnails"`
	PublishedAt    string                 `json:"published_at"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
}
