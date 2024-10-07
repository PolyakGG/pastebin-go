package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pastebin-go/internal/service"
)

type PasteHandler struct {
	service *service.PasteService
}

func NewPasteHandler(service *service.PasteService) *PasteHandler {
	return &PasteHandler{service: service}
}

func (h *PasteHandler) CreatePaste(c *gin.Context) {
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.CreatePaste(input.Title, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Метод для обрабатывания HTTP-запроса GET

// GetAllPastes retrieves all pastes
func (h *PasteHandler) GetAllPastes(c *gin.Context) {
	pastes, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve pastes"})
		return
	}
	c.JSON(http.StatusOK, pastes)
}
func (h *PasteHandler) GetPastePage(c *gin.Context) {
	id := c.Param("id")
	paste, err := h.service.GetPasteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paste not found"})
		return
	}

	c.HTML(http.StatusOK, "paste.html", gin.H{
		"title":   paste.Title,
		"content": paste.Content,
	})
}
