package main

import (
	"airbnb-scraper/config"
	"airbnb-scraper/services"
	"airbnb-scraper/storage"
	"airbnb-scraper/utils"

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

	// Show report
	report := utils.PropertyReport(properties)
	fmt.Println("\n---------- Report ----------")
	fmt.Println(report)
}
