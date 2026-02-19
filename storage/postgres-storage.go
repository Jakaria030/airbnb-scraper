package storage

import (
	"airbnb-scraper/config"
	"airbnb-scraper/models"
	"fmt"

	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB(connString string) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.TIMEOUT*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return err
	}

	if err := pool.Ping(ctx); err != nil {
		return err
	}

	Pool = pool
	log.Println("Connected to PostgreSQL")

	err = createTable()
	if err != nil {
		return err
	}

	createIndexesSQL := `
		CREATE INDEX IF NOT EXISTS idx_properties_price ON properties (price);
		CREATE INDEX IF NOT EXISTS idx_properties_location ON properties (location);
	`
	_, err = pool.Exec(ctx, createIndexesSQL)
	if err != nil {
		return err
	}

	fmt.Println("Table and indexes created successfully using pgxpool.")
	return nil
}

// createTable ensures table exists
func createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS properties (
		id SERIAL PRIMARY KEY,
		title TEXT,
		description TEXT,
		location TEXT,
		url TEXT UNIQUE,
		price DOUBLE PRECISION,
		rating DOUBLE PRECISION,
		created_at TIMESTAMP DEFAULT NOW()
	);
	`

	_, err := Pool.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

// Insert Properties
func InsertProperties(properties []models.Property) error {
	ctx := context.Background()
	batch := &pgx.Batch{}

	query := `
	INSERT INTO properties (title, description, location, url, price, rating)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (url) DO NOTHING;
	`

	for _, p := range properties {
		batch.Queue(query,
			p.Title,
			p.Description,
			p.Location,
			p.URL,
			p.Price,
			p.Rating,
		)
	}

	br := Pool.SendBatch(ctx, batch)
	defer br.Close()

	// Execute all queued queries
	for range properties {
		_, err := br.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}
