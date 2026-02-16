package tests

import (
	"testing"

	"github.com/Anryan2/URL_Shortener/service"
	"github.com/Anryan2/URL_Shortener/storage"
)

func TestMemory(t *testing.T) {
	store := storage.NewMemoryStorage()

	originalURL := "https://example.com"
	shortURL := service.GenerateHash(originalURL)

	store.Post(originalURL, shortURL)

	result := store.Get(shortURL)

	if result != originalURL {
		t.Errorf("Ожидался URL %s, получен %s", originalURL, result)
	}
}
