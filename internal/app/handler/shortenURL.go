package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortenURL(c *gin.Context) {
	body := c.Request.Body

	data, err := io.ReadAll(body)
	if err != nil {
		http.Error(c.Writer, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if len(data) == 0 {
		http.Error(c.Writer, "", http.StatusBadRequest)
		return
	}

	bodyStr := string(data)
	shortURL := h.services.URL.ShortenURL(bodyStr)

	c.String(http.StatusCreated, shortURL)
}
