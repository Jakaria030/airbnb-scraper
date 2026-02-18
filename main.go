package main

import (
	"airbnb-scraper/models"

	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	fmt.Println("Hello Airbnb")

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	url := "https://www.airbnb.com/s/Kuala-Lumpur/homes?place_id=ChIJ5-rvAcdJzDERfSgcL1uO2fQ&refinement_paths%5B%5D=%2Fhomes&flexible_trip_lengths%5B%5D=weekend_trip&date_picker_type=FLEXIBLE_DATES&search_type=HOMEPAGE_CAROUSEL_CLICK"

	var Properties []models.Property

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("div.c965t3n", chromedp.ByQuery),

		chromedp.Evaluate(`
			Array.from(document.querySelectorAll("div.c965t3n"))
				.map(card => {
					const baseUrl = "https://www.airbnb.com";

					// ----- Title -----
					const titleElement = card.querySelector("[data-testid='listing-card-title']");
					const title = titleElement ? titleElement.textContent.trim() : "";

					// ----- Description -----
					const descriptionElement = card.querySelector("[data-testid='listing-card-name']");
					description = descriptionElement ? descriptionElement.textContent.trim() : "";

					// ----- Location -----
					let location = "";
					if (titleElement) {
						const titleText = titleElement.textContent.trim();
						if (titleText.includes(" in ")) {
							location = titleText.split(" in ")[1];
						}
					}

					// ----- URL -----
					const urlElement = card.querySelector("a[href^='/rooms/']");
					let listingUrl = "";
					if (urlElement) {
						const href = urlElement.getAttribute("href").split("?")[0];
						listingUrl = baseUrl + href;
					}

					// ----- Price Per Night -----
					const priceRow = card.querySelector("[data-testid='price-availability-row']");
					let price = 0;
					if (priceRow) {
						const button = priceRow.querySelector("button");
						if (button) {
							const priceSpan = Array.from(button.querySelectorAll("span")).find(el => el.textContent.trim().startsWith("$"));

							if (priceSpan) {
								price = parseFloat(priceSpan.textContent.trim().replace("$", ""));
							}
						}

						const textContent = priceRow.innerText;
						const match = textContent.match(/(\d+)\s*nights?/);

						let nights = 0;
						if (match) {
							nights = parseInt(match[1]);
						}

						if (nights > 0) {
						    price = price / nights;
						}
					}

					// ----- Rating -----
					const ratingElement = Array.from(card.querySelectorAll("span[aria-hidden='true']")).find(el => /^\d+(\.\d+)?\s*\(\d+\)/.test(el.textContent.trim()));
					let rating = 0;
					if (ratingElement) {
						const match = ratingElement.textContent.trim().match(/^\d+(\.\d+)?/);
						if (match) {
							rating = parseFloat(match[0]);
						}
					}

					// ----- Return Fields -----
					return {
						Title: title,
						Description: description,
						Location: location,
						URL: listingUrl,
						Price: price,
						Rating: rating
					};
				});
		`, &Properties),
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, p := range Properties {
		fmt.Println("-----------", i, "---------------")
		fmt.Println("Title:", p.Title)
		fmt.Println("Description:", p.Description)
		fmt.Println("Location:", p.Location)
		fmt.Println("URL:", p.URL)
		fmt.Println("Price:", p.Price)
		fmt.Println("Rating:", p.Rating)
	}
}
