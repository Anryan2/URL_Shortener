package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Anryan2/URL_Shortener/service"
	"github.com/Anryan2/URL_Shortener/storage"

	_ "github.com/lib/pq"
)

var store storage.Storage

func main() {
	//Читаем значения из переменных окружения
	defaultStorage := os.Getenv("STORAGE_TYPE")
	defaultDbConn := os.Getenv("DB_CONN_STR")

	storageType := flag.String("storage", defaultStorage, "Тип хранилища: memory или postgres")
	dbConnStr := flag.String("db", defaultDbConn, "Строка подключения к PostgreSQL")
	flag.Parse()

	//Выбираем хранилище ссылок
	switch *storageType {
	case "postgres":
		fmt.Println("Используем хранилище:", *storageType)
		store = storage.NewDBStorage(*dbConnStr)
	case "memory":
		fmt.Println("Используем хранилище:", *storageType)
		store = storage.NewMemoryStorage()
	default:
		log.Fatal("Неизвестный тип хранилища")
	}

	//Создаем маршрутизатор
	mux := http.NewServeMux()

	mux.HandleFunc("/shorten", service.ShortenHandler(store))
	mux.HandleFunc("/", service.OriginalHandler(store))

	//Запускаем сервер
	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
