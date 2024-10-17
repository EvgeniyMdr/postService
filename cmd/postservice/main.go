package main

import (
	"fmt"
	"github.com/EvgeniyMdr/postService/db"
	httpInternal "github.com/EvgeniyMdr/postService/internal/http"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	database, err := db.ConnectToDB()

	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	defer func() {
		if err := database.Close(); err != nil {
			log.Fatalf("Ошибка закрытия соединения с базой данных: %v", err)
		}
	}()

	r := httpInternal.SetupRouter(database)

	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

	fmt.Println("Server started at port 8080")
}
