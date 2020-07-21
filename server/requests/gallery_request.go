package requests

type GalleryRequest struct {
	WebContentCategoryID string   `json:"web_content_category_id"`
	GalleryName          string   `json:"gallery_name"`
	ImagesID             []string `json:"images_id"`
}
