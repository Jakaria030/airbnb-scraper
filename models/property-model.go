package models

type Property struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	URL         string  `json:"url"`
	Price       int64   `json:"price"`
	Rating      float64 `json:"rating"`
}
