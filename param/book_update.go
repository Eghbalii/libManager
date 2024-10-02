package param

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"publish_date"`
}

type UpdateBookResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"publish_date"`
	Status      string `json:"status"`
}
