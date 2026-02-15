package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Anryan2/URL_Shortener/service"

	_ "github.com/lib/pq"
)

var storage storage.Storage

func main() {
	defaultStorage = os.Getenv("STORAGE_TYPE")

	var fullURL string
	fmt.Scanln(&fullURL)
	fmt.Println(fullURL)
	smallURL := service.GenerateHash(fullURL)
	fmt.Println(smallURL)

	switch *storageType {
	case "postgres":
		fmt.Println("Используем хранилище:", *storageType)
		storage = storage.NewPostgresStorage(*dbConnStr)
	case "memory":
		fmt.Println("Используем хранилище:", *storageType)
		storage = storage.NewMemoryStorage()
	default:
		log.Fatal("Неизвестный тип хранилища")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/shorten", service.ShortenHandler(storage))
	mux.HandleFunc("/", service.OriginalHandler(storage))

	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
