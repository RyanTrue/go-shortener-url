package storage

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/RyanTrue/go-shortener-url.git/internal/common/config"
	"net/url"
)

type urlService struct {
	repo   map[string]string
	config config.AppConfig
}

func (u urlService) ShortenURL(body string) string {
	hasher := hmac.New(sha256.New, []byte(u.config.Server.ServerAddr))
	_, err := hasher.Write([]byte(fmt.Sprintf("http://%s/api/v1/shorten_url?url=%s", u.config.Server.ServerAddr, url.QueryEscape(body))))
	if err != nil {
		return ""
	}
	h := hasher.Sum(nil)
	shortURL := ""
	for i := 0; i < 8; i++ {
		shortURL += string(h[i])
	}
	url, err := url.ParseRequestURI(shortURL)
	if err != nil {
		return ""
	}
	u.repo[url.String()] = body
	return url.String()
}

func (u *urlService) ExpandURL(path string) (string, error) {
	if value, ok := u.repo[path]; ok {
		return value, nil
	}
	return "", fmt.Errorf("URL path '%s' not found", path)
}
