package main

import (
	"airbnb-scraper/services"

	"fmt"
)

func main() {
	fmt.Println("Starting Airbnb Scraper...")

	properties, err := services.ScrapeAirbnb()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, p := range properties {
		fmt.Println("-----------", i, "---------------")
		fmt.Println("Title:", p.Title)
		fmt.Println("Description:", p.Description)
		fmt.Println("Location:", p.Location)
		fmt.Println("URL:", p.URL)
		fmt.Println("Price:", p.Price)
		fmt.Println("Rating:", p.Rating)
	}
}
