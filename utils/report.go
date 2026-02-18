package utils

import (
	"airbnb-scraper/models"
	"fmt"
	"sort"
)

func PropertyReport(properties []models.Property) string {
	totalListings := len(properties)

	// Calculate Average, Min, Max Price
	var totalPrice float64 = 0.0
	var minPrice float64 = properties[0].Price
	var maxPrice float64 = properties[0].Price
	var mostExpensiveProperty models.Property

	for _, property := range properties {
		totalPrice += property.Price
		if property.Price < minPrice {
			minPrice = property.Price
		}
		if property.Price > maxPrice {
			maxPrice = property.Price
			mostExpensiveProperty = property
		}
	}

	averagePrice := totalPrice / float64(totalListings)

	// Listings per Location
	locationCount := make(map[string]int)
	for _, property := range properties {
		locationCount[property.Location]++
	}

	// Top 5 Highest Rated Properties
	sort.Slice(properties, func(i, j int) bool {
		return properties[i].Rating > properties[j].Rating
	})

	top5Rated := make([]string, 0)
	for i := 0; i < 5 && i < len(properties); i++ {
		top5Rated = append(top5Rated, fmt.Sprintf("%d. %s — %.2f", i+1, properties[i].Title, properties[i].Rating))
	}

	// Prepare the report
	report := fmt.Sprintf(
		"Total Listings Scraped: %d\n\nAverage Price: %.2f\nMinimum Price: %.2f\nMaximum Price: %.2f\n\nMost Expensive Property:\nTitle: %s\nDescription: %s\nPrice: %.2f\nRating: %.2f\nURL: %s\nLocation: %s\n\nListings per Location:\n",
		totalListings, averagePrice, minPrice, maxPrice, mostExpensiveProperty.Title, mostExpensiveProperty.Description, mostExpensiveProperty.Price, mostExpensiveProperty.Rating, mostExpensiveProperty.URL, mostExpensiveProperty.Location,
	)

	// Add listings per location
	for location, count := range locationCount {
		report += fmt.Sprintf("%s: %d\n", location, count)
	}

	// Add top 5 highest rated properties
	report += "\nTop 5 Highest Rated Properties:\n"
	for _, rating := range top5Rated {
		report += rating + "\n"
	}

	return report
}
