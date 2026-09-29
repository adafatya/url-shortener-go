package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"url-shortener-go/internal/delivery/http/handler"
	"url-shortener-go/internal/helper"
	"url-shortener-go/internal/repository/memory"
	"url-shortener-go/internal/usecase"
)

func main() {
	dbHost := helper.EnvReader("DB_HOST", "localhost")
	dbPort := helper.EnvReader("DB_PORT", "5432")
	dbUser := helper.EnvReader("DB_USER", "root")
	dbPass := helper.EnvReader("DB_PASS", "")
	dbName := helper.EnvReader("DB_NAME", "my_app")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	urlRepository := memory.NewURLRepository()
	urlService := usecase.NewURLService(urlRepository)
	urlHandler := handler.NewURLHandler(urlService)

	port := helper.EnvReader("APP_PORT", "8080")
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
