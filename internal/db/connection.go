package db

import (
	"database/sql"
	"fmt"
	"github.com/EvgeniyMdr/postService/internal/config"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectToDB(dbConfig config.DbConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к базе данных: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к базе данных: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Успешное подключение к базе данных!")
	return db, nil
}
