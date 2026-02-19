package main

import (
	"airbnb-scraper/config"
	"airbnb-scraper/services"
	"airbnb-scraper/storage"
	"airbnb-scraper/utils"

	"fmt"
)

func main() {
	// Connect to the PostgreSQL database
	storage.InitDB(config.CONNECTION_STRING)

	fmt.Println("Starting Airbnb Scraper...")

	// Call base scraper
	properties, err := services.ScrapeAirbnb()
	if err != nil {
		fmt.Println("Error to Scrape:", err)
		return
	}
	fmt.Println("\nAirbnb data fetch successfully.")

	// Save into csv file
	err = storage.SavePropertiesToCSV(properties, config.FILE_PATH)
	if err != nil {
		fmt.Println("Error To Save:", err)
		return
	}
	fmt.Println("Airbnb data saved into csv file successfully.")

	// Save property to the database
	err = storage.InsertProperties(properties)
	if err != nil {
		fmt.Println("Airbnb properties is not saved into postgresql")
	}else{
		fmt.Println("Airbnb data saved into postgresql database successfully.")
	}

	// Show report
	report := utils.PropertyReport(properties)
	fmt.Println("\n---------- Report ----------")
	fmt.Println(report)
}
