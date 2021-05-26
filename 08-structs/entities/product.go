package entities

import "fmt"

type Product struct {
	Id        int
	Name      string
	Cost      float32
	Units     int
	Category  string
	IsInStock bool
}

func NewProduct(id int, name string, cost float32, units int, category string, isInStock bool) *Product {
	return &Product{id, name, cost, units, category, isInStock}
}

func (productPtr *Product) ApplyDiscount(discount float32) {
	productPtr.Cost = productPtr.Cost * ((100 - discount) / 100)
}

func (productPtr *Product) Print() {
	fmt.Println("Product - ", productPtr)
}

func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, cost = %f, units = %d, category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}
