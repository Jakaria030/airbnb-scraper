package services

import (
	"airbnb-scraper/config"
	"airbnb-scraper/models"
	"airbnb-scraper/scraper"
	"fmt"

	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func ScrapeAirbnb(searchURLS []string) ([]models.Property, error) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, config.TIMEOUT*time.Second)
	defer cancel()

	var allProperties []models.Property

	for indx, url := range searchURLS {
		fmt.Printf("Page No: %d scraping...\n", indx+1)

		properties, err := scraper.ScrapePage(ctx, url)

		if err != nil {
			fmt.Printf("Error while scraping page No: %d\n", indx+1)
			return nil, err
		}

		allProperties = append(allProperties, properties...)
	}

	return allProperties, nil
}
