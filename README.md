# 🏡 Airbnb Scraper 
A web scraper built with Go and Chromedp to extract property data from Airbnb.


## 📂 Project Structure
```bash
airbnb-scraper
│──── config/  
│       └── config.go
├──── data/
│       └── properties.csv
├──── models/
│       └── property-model.go
├──── scraper/
│       └── scraper.go
├──── services/
│       └── scraper-service.go
├──── storage/
│       └── csv-storage.go
│       └── postgres-storage.go
├──── utils/
│       └── helper.go
│       └── report.go
├──── .gitignore
├──── go.mod
├──── go.sum
├──── main.go
└──── README.md  
```


## 🚀 Getting Started
Follow these steps to set up and run the project locally:

1. Clone the repository:
    ```bash
    git clone https://github.com/Jakaria030/airbnb-scraper.git
    ```
2. Navigate to the project directory:
    ```bash
    cd airbnb-scraper
    ```
3. Install Dependencies
    ```bash
    go mod tidy
    ```
4. Install PostgreSQL (Optional)
    ```bash
    docker run -d --name scraper-postgres -e POSTGRES_DB=scraperdb -e POSTGRES_USER=scraperuser -e POSTGRES_PASSWORD=scraperpass -p 5432:5432 -v scraper_pgdata:/var/lib/postgresql/data --restart unless-stopped postgres:16
    ```
5. Run the Scraper
    ```bash
    go run main.go
    ```

## Additional Resources
- [Go Documentation](https://go.dev/doc/)
- [Chromedp Documentation](https://pkg.go.dev/github.com/chromedp/chromedp)
- [Docker Documentation](https://docs.docker.com/get-started/)
