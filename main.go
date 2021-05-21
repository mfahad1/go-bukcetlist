package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mfahad1/go-bukcetlist/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api-", log.LstdFlags)

	productHandler := handlers.NewProduct(l)

	muxRouter := mux.NewRouter()

	getRouter := muxRouter.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/product", productHandler.GetProducts)

	postRouter := muxRouter.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/product", productHandler.AddProduct)

	putRouter := muxRouter.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/product/{id:[0-9]+}", productHandler.UpdateProduct)

	server := &http.Server{
		Addr:	":9090",
		Handler: muxRouter,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}


	server.ListenAndServe()
}