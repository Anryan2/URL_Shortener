package storage_test

import (
	"testing"

	"github.com/Anryan2/URL_Shortener/internal/service"
	"github.com/Anryan2/URL_Shortener/internal/storage"
)

func TestMemory(t *testing.T) {
	store := storage.NewMemoryStorage()

	originalURL := "https://example.com"
	shortURL := service.GenerateHash(originalURL)

	store.Insert(originalURL, shortURL)

	result := store.Get(shortURL)

	if result != originalURL {
		t.Errorf("Ожидался URL %s, получен %s", originalURL, result)
	}
}
