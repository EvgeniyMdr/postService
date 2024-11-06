package main

import (
	"fmt"
	"github.com/EvgeniyMdr/postService/internal/config"
	"github.com/EvgeniyMdr/postService/internal/db"
	httpInternal "github.com/EvgeniyMdr/postService/internal/http"
	"github.com/EvgeniyMdr/postService/internal/repositories"
	"github.com/EvgeniyMdr/postService/internal/services"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	mainConfig := config.NewServiceConfig()
	database, err := db.ConnectToDB(mainConfig.GetDbSettings())

	postRepo := repositories.NewPostRepository(database)

	postService := services.NewPostService(postRepo)

	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	defer func() {
		if err := database.Close(); err != nil {
			log.Fatalf("Ошибка закрытия соединения с базой данных: %v", err)
		}
	}()

	r := httpInternal.SetupRouter(postService)
	httpSettings := mainConfig.GetHttpSettings()
	address := httpSettings.GetAddress()

	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

	fmt.Println("Server started at port 8080")
}
