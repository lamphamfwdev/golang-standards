package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/config"
)

func NewPostgres(cfg config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("cannot connect database:", err)
	}

	fmt.Println("Connected PostgreSQL")

	return db
}
