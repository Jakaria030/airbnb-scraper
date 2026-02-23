#!/bin/bash

set -e  # Exit immediately if a command fails

echo "🚀 Starting Airbnb Scraper..."

# -----------------------------
# Check Go installation
# -----------------------------
if ! command -v go &> /dev/null
then
    echo "❌ Go is not installed."
    echo "👉 Please install Go: https://go.dev/dl/"
    exit 1
fi

echo "✅ Go found: $(go version)"

# -----------------------------
# Install Go dependencies
# -----------------------------
echo "📦 Installing Go dependencies..."
go mod tidy

# -----------------------------
# Check Docker (Optional DB Setup)
# -----------------------------
if command -v docker &> /dev/null
then
    echo "🐳 Docker detected"

    if [ "$(docker ps -aq -f name=scraper-postgres)" ]; then
        echo "📦 PostgreSQL container exists. Starting..."
        docker start scraper-postgres > /dev/null || true
    else
        echo "🛠 Creating PostgreSQL container..."
        docker run -d \
            --name scraper-postgres \
            -e POSTGRES_DB=scraperdb \
            -e POSTGRES_USER=scraperuser \
            -e POSTGRES_PASSWORD=scraperpass \
            -p 5432:5432 \
            -v scraper_pgdata:/var/lib/postgresql/data \
            --restart unless-stopped \
            postgres:16
    fi
else
    echo "⚠️ Docker not installed. Skipping PostgreSQL setup."
fi

# -----------------------------
# Run the scraper
# -----------------------------
echo "🔥 Running scraper..."
go run main.go