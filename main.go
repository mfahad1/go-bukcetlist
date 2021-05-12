package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mfahad1/go-bukcetlist/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api-", log.LstdFlags)

	productHandler := handlers.NewProduct(l)

	serverHandler := http.NewServeMux()
	serverHandler.Handle("/", productHandler)

	server := &http.Server{
		Addr:	":9090",
		Handler: serverHandler,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}


	server.ListenAndServe()
}