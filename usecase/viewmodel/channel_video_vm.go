package viewmodel

type ChannelVideoVm struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	ChannelName string                 `json:"channel_name"`
	Thumbnails  map[string]interface{} `json:"thumbnails"`
	Description string                 `json:"description"`
	PublishedAt string                 `json:"published_at"`
}
