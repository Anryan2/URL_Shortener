package service_test

import (
	"testing"
	"github.com/Anryan2/URL_Shortener/internal/service"
)

func TestGenerateHash(t *testing.T) {
	url := "http://example.com"
	shortUrl := service.GenerateHash(url)
	if len(shortUrl) != 10 {
		t.Errorf("Ожидалась короткая ссылка дляной 10 символов, получена строка дляной %d символов", len((shortUrl)))
	}

	secondURL := service.GenerateHash(url)
	if shortUrl != secondURL {
		t.Errorf("Строки различны %s не равен %s", shortUrl, secondURL)
	}
}
