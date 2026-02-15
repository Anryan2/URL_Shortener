package storage

import "sync"

type Storage interface {
	Post(originalURL, shortURL string)
	Get(shortURL string) string
	checkExists(originalURL string) bool
}

type MemoryStorage struct {
	URLs map[string]string
	mu   sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		URLs: make(map[string]string),
	}
}

func (m *MemoryStorage) Post(originalURL, shortURL string) {
	if !m.checkExists(shortURL) {
		m.mu.Lock()
		defer m.mu.Unlock()
		m.URLs[shortURL] = originalURL
	}

}

func (m *MemoryStorage) Get(shortURL string) string {
	if !m.checkExists(shortURL) {
		return ""
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.URLs[shortURL]
}

func (m *MemoryStorage) checkExists(shortURL string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, exist := m.URLs[shortURL]
	return exist
}
