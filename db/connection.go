package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	if dbPort == "" {
		dbPort = "5432"
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к базе данных: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к базе данных: %v", err)
	}

	// Установка параметров пула соединений
	db.SetMaxOpenConns(25)                 // Максимум открытых соединений
	db.SetMaxIdleConns(25)                 // Максимум "простаивающих" соединений
	db.SetConnMaxLifetime(5 * time.Minute) // Максимальное время жизни соединения

	log.Println("Успешное подключение к базе данных!")
	return db, nil
}
