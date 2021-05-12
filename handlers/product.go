package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/mfahad1/go-bukcetlist/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct (l*log.Logger) *Products {
	return &Products{l}
}

func (p*Products) ServeHTTP(rw http.ResponseWriter, req *http.Request ) {
	if req.Method == http.MethodGet {
		p.getProducts(rw);

		return
	}

	if req.Method == http.MethodPost {
		p.addProduct(rw, req);

		return
	}

	if req.Method == http.MethodPut {
		r := regexp.MustCompile("([0-9]+)")
		g := r.FindAllString(req.URL.RequestURI(), -1)
		p.l.Println(g)
		if len(g) != 1 || len(g[0]) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)

			return
		}

		idStr := g[0]

		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(rw, "Wrong ID passed", http.StatusBadGateway)
		}

		p.updateProduct(rw, req, id)

	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p*Products) getProducts(rw http.ResponseWriter) {
	listOfProducts := data.GetProducts()

	err := listOfProducts.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to Marshall json", http.StatusInternalServerError)
	}
}

func (p*Products) addProduct(rw http.ResponseWriter, req*http.Request) {
	product := &data.Product{}

	err := product.FromJson(req.Body)

	if err != nil {
		http.Error(rw, "Unable to unMarshall json", http.StatusInternalServerError)
	}

	data.AddProduct(product)
}

func (p*Products) updateProduct(rw http.ResponseWriter, req *http.Request, id int) {
	prod := &data.Product{};
	parseError := prod.FromJson(req.Body);

	if parseError != nil {
		http.Error(rw, "Not able to parse request", http.StatusBadRequest)
	}

	data.UpdateProduct(prod, id)
}