package storage_test

import (
	"testing"

	"github.com/Anryan2/URL_Shortener/internal/service"
	"github.com/Anryan2/URL_Shortener/internal/storage"
	_ "github.com/lib/pq"
)

func TestDB(t *testing.T) {
	connStr := "postgres://postgres:12345@localhost/postgres?sslmode=disable"
	store := storage.NewDBStorage(connStr)

	originalURL := "https://example.com"
	shortURL := service.GenerateHash(originalURL)

	store.Insert(originalURL, shortURL)

	result := store.Get(shortURL)

	if result != originalURL {
		t.Errorf("Ожидался URL %s, получен %s", originalURL, result)
	}

}
