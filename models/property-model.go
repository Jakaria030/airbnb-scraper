package models

type Property struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	URL         string  `json:"url"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
}

type ScrapeJob struct {
	SectionIndx int    `json:"sectionIndx"`
	Page        int    `json:"page"`
	Url         string `json:"url"`
}

type Result struct {
	Properties []Property `json:"properties"`
	Err        error      `json:"err"`
	Job        ScrapeJob  `json:"job"`
}
