package main

import (
	"log"
	"mygame/package/game"
	"net/http"
)

func main() {
	http.HandleFunc("/", game.Handler)
	log.Println("Сервер запущен на порту :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

