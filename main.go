package main

import (
	"fmt"

	"github.com/Anryan2/URL_Shortener/service"

	_ "github.com/lib/pq"
)

func main() {

	var fullURL string
	fmt.Scanln(&fullURL)
	fmt.Println(fullURL)
	smallURL := service.GenerateHash(fullURL)
	fmt.Println(smallURL)

}
