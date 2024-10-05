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
	// Раздача статических файлов
	router.Static("/static", "./static")
	// Маршрут для основной страницы.
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	// API для создания и получения pastes
	router.POST("/pastes", pasteHandler.CreatePaste)
	router.GET("pastes/:id", pasteHandler.GetPaste)
	router.GET("/pastes", pasteHandler.GetAllPastes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
