package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mfahad1/go-bukcetlist/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct (l*log.Logger) *Products {
	return &Products{l}
}

func (p*Products) GetProducts(rw http.ResponseWriter, req*http.Request) {
	listOfProducts := data.GetProducts()

	err := listOfProducts.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to Marshall json", http.StatusInternalServerError)
	}
}

func (p*Products) AddProduct(rw http.ResponseWriter, req*http.Request) {
	product := &data.Product{}

	err := product.FromJson(req.Body)

	if err != nil {
		http.Error(rw, "Unable to unMarshall json", http.StatusInternalServerError)
	}

	data.AddProduct(product)
}

func (p*Products) UpdateProduct(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)

		return
	}

	prod := &data.Product{};
	parseError := prod.FromJson(req.Body);

	if parseError != nil {
		http.Error(rw, "Not able to parse request", http.StatusBadRequest)
	}

	data.UpdateProduct(prod, id)
}

type KeyProduct struct {}

func (p*Products) MiddlewareProductValidation(next http.Handler) http.Handler{
	return http.HandlerFunc(func (rw http.ResponseWriter, req *http.Request) {
		prod := &data.Product{};
		parseError := prod.FromJson(req.Body);

		if parseError != nil {
			http.Error(rw, "Not able to parse request", http.StatusBadRequest)

			return
		}

		ctx := context.WithValue(KeyProduct{}, prod)
	})
}