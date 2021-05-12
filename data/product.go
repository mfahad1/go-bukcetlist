package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)

	return e.Decode(p)
}

var timeNow = time.Now().UTC().String()

var productList = []*Product{
	{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milk coffee",
		Price: 2.45,
		SKU: "abc23",
		CreatedOn: timeNow,
		UpdatedOn: timeNow,
	},
	{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee without milk",
		Price: 1.99,
		SKU: "zxxaaa",
		CreatedOn: timeNow,
		UpdatedOn: timeNow,
	},
}

func GetProducts () Products{
	return productList;
}

func AddProduct (p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)

}

func UpdateProduct (p *Product, id int) {
	for count := range productList {
		if productList[count].ID == id {
			productList[count] = p
		}
	}
}

func getNextID () int {
	return len(productList)
}