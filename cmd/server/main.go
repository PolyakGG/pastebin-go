package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"pastebin-go/internal/handlers"
	"pastebin-go/internal/repository"
	"pastebin-go/internal/service"
)

func main() {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "./pastebin.db"
	}
	repo, err := repository.NewPasteRepository(dbPath)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	svc := service.NewPasteService(repo)
	pasteHandler := handlers.NewPasteHandler(svc)
	router := gin.Default()

	router.POST("/pastes", pasteHandler.CreatePaste)
	router.GET("pastes/:id", pasteHandler.GetPaste)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
