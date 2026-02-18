package services

import (
	"airbnb-scraper/config"
	"airbnb-scraper/models"
	"airbnb-scraper/scraper"

	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func ScrapeAirbnb() ([]models.Property, error) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, config.TIMEOUT*time.Second)
	defer cancel()

	return scraper.ScrapePage(ctx, config.SEARCH_URL)
}
