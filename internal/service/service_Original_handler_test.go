package service_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Anryan2/URL_Shortener/internal/service"
	"github.com/Anryan2/URL_Shortener/internal/storage"
)

func TestOriginalHandler(t *testing.T) {
	store := storage.NewMemoryStorage()
	originalURL := "https://example.com"
	shortURl := service.GenerateHash(originalURL)

	store.Insert(originalURL, shortURl)

	req := httptest.NewRequest("GET", "/"+shortURl, nil)
	rec := httptest.NewRecorder()

	handler := service.OriginalHandler(store)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusFound {
		t.Errorf("Ожидался код статуса 302, но получен %d", rec.Code)
	}
	gettenURL := rec.Header().Get("location")
	if gettenURL != originalURL {
		t.Errorf("Ожидался URl %s, но получен %s", originalURL, gettenURL)
	}
}
