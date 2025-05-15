package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/noBthd/5x30-fitness-api/internal/config"
)

var (
	DB *sql.DB
)

func ConnectDB(cfg *config.Config) {
	dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.SSLMode,
    )
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB = database
}