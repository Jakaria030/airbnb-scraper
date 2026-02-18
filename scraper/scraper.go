package scraper

import (
	"airbnb-scraper/config"
	"airbnb-scraper/models"

	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func ScrapePage(ctx context.Context, url string) ([]models.Property, error) {

	var properties []models.Property

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("div.c965t3n", chromedp.ByQuery),

		chromedp.Evaluate(fmt.Sprintf(`
		Array.from(document.querySelectorAll("div.c965t3n"))
			.map(card => {

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
					listingUrl = "%s" + href;
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
	`, config.BASE_URL), &properties),
	)

	return properties, err
}
