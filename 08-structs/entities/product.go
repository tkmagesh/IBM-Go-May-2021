package entities

import "fmt"

type Product struct {
	Id        int
	Name      string
	Cost      float32
	Units     int
	Category  string
	IsInStock bool
	dummy     string
}

func NewProduct(id int, name string, cost float32, units int, category string, isInStock bool) *Product {
	return &Product{id, name, cost, units, category, isInStock, "Dummy string"}
}

func (productPtr *Product) ApplyDiscount(discount float32) {
	productPtr.Cost = productPtr.Cost * ((100 - discount) / 100)
}

func (productPtr *Product) Print() {
	fmt.Println("Product - ", productPtr)
}
