package storage

import (
	"airbnb-scraper/models"

	"encoding/csv"
	"fmt"
	"os"
)

func SavePropertiesToCSV(properties []models.Property, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create file %s: %v", filePath, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Title", "Description", "Location", "URL", "Price", "Rating"})
	for _, property := range properties {
		record := []string{
			property.Title,
			property.Description,
			property.Location,
			property.URL,
			fmt.Sprintf("%.2f", property.Price),
			fmt.Sprintf("%.2f", property.Rating),
		}
		err := writer.Write(record)
		if err != nil {
			return fmt.Errorf("could not write property to CSV: %v", err)
		}
	}

	return nil
}
