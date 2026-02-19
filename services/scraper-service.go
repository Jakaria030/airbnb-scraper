package services

import (
	"airbnb-scraper/config"
	"airbnb-scraper/models"
	"airbnb-scraper/scraper"
	"airbnb-scraper/utils"

	"fmt"
	"log"

	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func ScrapeAirbnb() ([]models.Property, error) {

	// context allocator
	allocCtx, _ := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("enable-automation", false),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/120 Safari/537.36"),
	)

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, config.TIMEOUT*time.Second)
	defer cancel()

	// Collect all links from home page
	var sectionLinks []string
	err := chromedp.Run(ctx,
		chromedp.Navigate(config.BASE_URL),
		chromedp.Sleep(utils.RandomDelay()),
		chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight)`, nil),
		chromedp.Sleep(utils.RandomDelay()),
		chromedp.Evaluate(`
			Array.from(document.querySelectorAll("a[href^='/s/']"))
			.map(a => a.href)
		`, &sectionLinks),
	)

	if err != nil {
		log.Fatal("Navigation error:", err)
	}

	// Deduplicate section links
	sectionLinks = utils.Unique(sectionLinks)

	if len(sectionLinks) == 0 {
		log.Fatal("No home page sections links found")
	}

	// Only first MAX_HOME_PAGE_LINK sections
	if len(sectionLinks) > config.MAX_HOME_PAGE_LINK {
		sectionLinks = sectionLinks[:config.MAX_HOME_PAGE_LINK]
	}

	// Loop sections
	var allProperties []models.Property
	for indx, sectionURL := range sectionLinks {
		fmt.Println("\nSECTION:", indx+1)

		for page := 1; page <= config.MAX_PAGE_NO; page++ {

			offset := (page - 1) * 20
			pageURL := utils.AddOffset(sectionURL, offset)

			fmt.Println("Visiting page no:", page)

			var properties []models.Property
			properties, err := scraper.ScrapePage(ctx, pageURL)

			if err != nil {
				log.Println("Scrape error:", err)
				continue
			} else {
				fmt.Println("Date fetched for page no:", page)
			}

			allProperties = append(allProperties, properties...)
		}
	}

	return allProperties, nil
}
