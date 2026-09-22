package main

import (
	"log"
	"net/http"
	"os"

	"url-shortener-go/internal/delivery/http/handler"
	"url-shortener-go/internal/repository/memory"
	"url-shortener-go/internal/usecase"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	urlRepository := memory.NewURLRepository()
	urlService := usecase.NewURLService(urlRepository)
	urlHandler := handler.NewURLHandler(urlService)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: urlHandler.Routes(),
	}
	log.Printf("starting app")
	log.Printf("url shortener listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
