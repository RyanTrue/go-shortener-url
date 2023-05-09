package service

type URL interface {
	ShortenURL(body string) string
	GetOriginalURL(path string) (string, error)
}

type Service struct {
	URL
}

func NewServiceURL(repo map[string]string) *Service {
	return &Service{
		URL: NewURLservice(),
	}
}
