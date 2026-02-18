package main

import (
	"airbnb-scraper/config"
	"airbnb-scraper/scraper"

	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	fmt.Println("Hello Airbnb")

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, config.TIMEOUT*time.Second)
	defer cancel()

	properties, err := scraper.ScrapePage(ctx, config.SEARCH_URL)

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
