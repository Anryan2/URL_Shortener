package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Anryan2/URL_Shortener/service"
	"github.com/Anryan2/URL_Shortener/storage"
)

func TestShortenHandler(t *testing.T) {
	store := storage.NewMemoryStorage()

	reqData := map[string]string{"url": "https://example.com"}
	reqBody, _ := json.Marshal(reqData)

	r := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewReader(reqBody))
	rec := httptest.NewRecorder()

	handler := service.ShortenHandler(store)
	handler.ServeHTTP(rec, r)

	if rec.Code != http.StatusOK {
		t.Errorf("Ожидался ответ 200, а получен %d", rec.Code)

	}

	var responce map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &responce)

	if err != nil {
		t.Fatalf("Не удалось демаршалировать ответ: %v", err)
	}
	if responce["short_url"] == "" {

		t.Errorf("Ожидалась короткая ссылка, получен пустой ответ")
	}

}
