package storage

import (
	"crypto/sha256"
	"fmt"
	"github.com/RyanTrue/go-shortener-url.git/internal/common/config"
)

type urlService struct {
	repo   map[string]string
	config config.AppConfig
}

func (u *urlService) ShortenURL(body string) string {
	hasher := sha256.New()                                              // используем SHA256 для безопасности
	hasher.Write([]byte(body))                                          // хешируем тело запроса
	hash := fmt.Sprintf("%x", hasher.Sum(nil)[:8])                      // преобразуем хеш в шестнадцатеричное представление
	shortURL := fmt.Sprintf("%s/%s", u.config.Server.DefaultAddr, hash) // создаем короткий URL

	// проверяем, существует ли уже созданный короткий URL для данного тела запроса
	if _, ok := u.repo[hash]; ok {
		return ""
	}

	u.repo[hash] = body // обновляем значение в хранилище
	return shortURL     // возвращаем короткий URL без ошибок
}

func (u *urlService) ExpandURL(path string) (string, error) {
	if value, ok := u.repo[path]; ok {
		return value, nil
	}
	return "", fmt.Errorf("URL path '%s' not found", path)
}
