package main

import (
	"airbnb-scraper/config"
	"airbnb-scraper/services"
	"airbnb-scraper/storage"

	"fmt"
)

func main() {
	fmt.Println("Starting Airbnb Scraper...")

	// Call base scraper
	properties, err := services.ScrapeAirbnb()
	if err != nil {
		fmt.Println("Error To Srape:", err)
		return
	}

	fmt.Println("Airbnb data fetch successfully.")

	// Save into csv file
	err = storage.SavePropertiesToCSV(properties, config.FILE_PATH)
	if err != nil {
		fmt.Println("Error To Save:", err)
		return
	}

	fmt.Println("Airbnb data saved into csv file successfully.")



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
