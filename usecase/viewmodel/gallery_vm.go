package viewmodel

type GalleryVm struct {
	ID          string `json:"id"`
	GalleryName string `json:"gallery_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type GalleryDetailVm struct {
	ID          string            `json:"id"`
	GalleryName string            `json:"gallery_name"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	DeletedAt   string            `json:"deleted_at"`
	Images      []GalleryImagesVm `json:"images"`
}
