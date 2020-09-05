package viewmodel

type VideoContentVm struct {
	ID        string `json:"id"`
	Channel   string `json:"channel"`
	ChannelID string `json:"channel_id"`
	Link      string `json:"link"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
