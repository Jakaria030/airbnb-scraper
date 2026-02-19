package services

import (
	"airbnb-scraper/config"
	"airbnb-scraper/models"
	"airbnb-scraper/scraper"
	"airbnb-scraper/utils"
	"fmt"
	"sync"

	"log"

	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func ScrapeAirbnb() ([]models.Property, error) {

	// context allocator
	allocCtx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("enable-automation", false),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/120 Safari/537.36"),
	)
	defer cancel()

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

	jobs := make(chan models.ScrapeJob)
	go func() {
		defer close(jobs)

		for indx, sectionURL := range sectionLinks {
			for page := 1; page <= config.MAX_PAGE_NO; page++ {

				offset := (page - 1) * 20
				jobs <- models.ScrapeJob{
					SectionIndx: indx + 1,
					Page:        page,
					Url:         utils.AddOffset(sectionURL, offset),
				}
			}
		}
	}()

	// Collect result
	results := make(chan models.Result)
	var wg sync.WaitGroup

	for worker := 0; worker < config.MAX_WORKERS; worker++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for job := range jobs {
				workerCtx, cancel := chromedp.NewContext(allocCtx)
				defer cancel()
				workerCtx, cancel = context.WithTimeout(workerCtx, config.TIMEOUT*time.Second)
				defer cancel()

				props, err := scraper.ScrapePage(workerCtx, job.Url)
				results <- models.Result{Properties: props, Err: err, Job: job}
			}

		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		ticker := time.NewTicker(10 * time.Second) // Change interval as needed
		defer ticker.Stop()

		for range ticker.C {
			// Print a random message every interval
			fmt.Println(utils.GetRandomMessage())
		}
	}()

	seen := sync.Map{}
	var allProperties []models.Property
	for result := range results {
		if result.Err != nil {
			log.Printf("Scrape error (section %d, page %d): %v", result.Job.SectionIndx, result.Job.Page, result.Err)
			continue
		}
		for _, p := range result.Properties {
			if _, loaded := seen.LoadOrStore(p.URL, struct{}{}); !loaded {
				allProperties = append(allProperties, p)
			}
		}
	}

	return allProperties, nil
}
