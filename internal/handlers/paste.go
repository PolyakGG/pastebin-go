package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pastebin-go/internal/service"
	"strconv"
)

type PasteHandler struct {
	service *service.PasteService
}

func NewPasteHandler(service *service.PasteService) *PasteHandler {
	return &PasteHandler{service: service}
}
func (h *PasteHandler) CreatePaste(c *gin.Context) {
	var input struct {
		Title   string `json:"title" binding:"required"'`
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

func (h *PasteHandler) GetPaste(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	paste, err := h.service.GetPaste(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paste not found"})
		return
	}
	c.JSON(http.StatusOK, paste)
}
