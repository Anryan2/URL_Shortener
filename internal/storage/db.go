package storage

import (
	"database/sql"
	"log"
)

// DBStorage - структура для хранения ссылок в базе данных, наследующая интерфейс Storage.
type DBStorage struct {
	db *sql.DB
}

// Функция NewDBStorage создает новое подключение к PostgreSQL.
func NewDBStorage(connection string) *DBStorage {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal("Ошибка при открытии базы данных:", err)
	}

	//Проверяем соединение.
	err = db.Ping()
	if err != nil {
		log.Fatal(connection, err)
	}

	//Создаем таблицу, если ее нет.
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id SERIAL PRIMARY KEY,
			short_url VARCHAR(10) UNIQUE NOT NULL,
			original_url TEXT UNIQUE NOT NULL
		);
		`)
	if err != nil {
		log.Fatal("Ошибка при создании таблицы:", err)
	}
	return &DBStorage{db: db}
}

func (db *DBStorage) Insert(originalURL, shortURL string) {
	_, err := db.db.Exec("INSERT INTO urls (short_url, original_url) VALUES ($1, $2)", shortURL, originalURL)
	if err != nil {
		log.Fatal("Ошибка метода Insert:", err)
	}

}

func (db *DBStorage) Get(shortURL string) string {
	var originalURL string
	err := db.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		log.Fatal("Ошибка метода Get:", err)
	}
	return originalURL
}

func (db *DBStorage) CheckExists(shortURL, originalURL string) bool {
	var newURL string
	err := db.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&newURL)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal("Ошибка метода CheckExists:", err)
	}
	return err == sql.ErrNoRows || newURL == originalURL
}
