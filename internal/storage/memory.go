package storage

import "sync"

//MemoryStorage - структура для хранения ссылок в памяти, наследующая интерфейс Storage.
type MemoryStorage struct {
	URLs map[string]string //URL[сокращенная ссылка] - полная ссылка
	mu   sync.RWMutex
}

//NewMemoryStorage создает новое хранилище в памяти
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		URLs: make(map[string]string),
	}
}

func (m *MemoryStorage) Insert(originalURL, shortURL string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.URLs[shortURL] = originalURL

}

func (m *MemoryStorage) Get(shortURL string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.URLs[shortURL]
}

func (m *MemoryStorage) CheckExists(shortURL, originalURL string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exist := m.URLs[shortURL]
	return !exist || value == originalURL
}
