package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortenURL(c *gin.Context) {
	body := c.Request.Body

	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			fmt.Printf("Failed to close body: %v", err)
		}
	}(body)

	data, err := io.ReadAll(body)
	if err != nil {
		http.Error(c.Writer, "Error reading request body", http.StatusInternalServerError)
		return
	}

	bodyStr := string(data)
	shortURL := h.services.URL.ShortenURL(bodyStr)

	c.String(http.StatusCreated, shortURL)
}

func (h *Handler) GetOriginalURL(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		http.Error(c.Writer, "Error reading id param", http.StatusInternalServerError)
		return
	}

	value, err := h.services.URL.GetOriginalURL(id)
	if err != nil {
		http.Error(c.Writer, "No original URL found", http.StatusNotFound)
		return
	}
	c.Header("Location", value)
	c.Status(http.StatusTemporaryRedirect)
}
