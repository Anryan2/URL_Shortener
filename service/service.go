package service

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Anryan2/URL_Shortener/storage"
)

// Структуры для входного json и для ответа
type shortenRequest struct {
	URL string `json:"url"`
}
type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

// обрабатывает создание короткой ссылки
func ShortenHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
			return
		}
		var request shortenRequest
		err = json.Unmarshal(body, &request)
		if err != nil || request.URL == "" {
			http.Error(w, "Неверный запрос", http.StatusBadRequest)
			return
		}

		shortURL := GenerateHash(request.URL)
		storage.Post(request.URL, shortURL)
		response := shortenResponse{ShortURL: "http://localhost:8080/" + shortURL}
		responseData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Ошибка формирования ответа", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(responseData); err != nil {

			http.Error(w, "Ошибка записи в ответ", http.StatusInternalServerError)
			return
		}
		return

	}
}

// Обрабатывает возврат полной ссылки по короткой
func OriginalHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		shortURL := strings.TrimPrefix(r.URL.Path, "/")
		if shortURL == "" {
			http.Error(w, "Короткая ссылка не найдена", http.StatusBadRequest)
			return
		}
		originalURL := storage.Get(shortURL)
		if originalURL == "" {
			http.Error(w, shortURL, http.StatusNotFound)
			return
		}

		http.Redirect(w, r, originalURL, http.StatusFound)
	}
}

// генерирует хэш из 10 символов
func GenerateHash(full string) string {
	sum := sha256.Sum224([]byte(full))
	encoded := base64.URLEncoding.EncodeToString(sum[:])
	return strings.ReplaceAll(strings.TrimRight(encoded, "="), "-", "_")[:10]
}
