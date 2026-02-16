package storage

//Storage - интерфейс, определяющий методы для работы с хранением URL.
type Storage interface {
	Insert(originalURL, shortURL string) //Сохраняет ссылку в памяти
	Get(shortURL string) string          //Вохвращает полную ссылку
}
