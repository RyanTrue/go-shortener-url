package storage

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/RyanTrue/go-shortener-url.git/internal/common/config"
)

type urlService struct {
	repo   map[string]string
	config config.AppConfig
}

func (u *urlService) ShortenURL(body string) string {
	sha1 := md5.Sum([]byte(body))
	hash := fmt.Sprintf("sha1-%s", hex.EncodeToString(sha1[:]))
	shortURL := fmt.Sprintf("%s/short/%s", u.config.Server.DefaultAddr, hash)
	if _, ok := u.repo[hash]; !ok {
		u.repo[hash] = body
	}
	return shortURL
}

func (u *urlService) ExpandURL(path string) (string, error) {
	if value, ok := u.repo[path]; ok {
		return value, nil
	}
	return "", fmt.Errorf("URL path '%s' not found", path)
}
