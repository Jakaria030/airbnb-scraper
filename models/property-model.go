package models

type Property struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	URL         string  `json:"url"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
}
