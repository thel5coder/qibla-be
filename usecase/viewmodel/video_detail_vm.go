package viewmodel

type VideoDetailVm struct {
	Title          string `json:"title"`
	ChannelName    string `json:"channel_name"`
	Description    string `json:"description"`
	EmbeddedPlayer string `json:"embedded_player"`
	PublishedAt    string `json:"published_at"`
}
