package entities

import "fmt"

type ProductList []Product

var x = 100

func (productList *ProductList) AddProduct(p *Product) {
	fmt.Println(p.dummy)
	*productList = append(*productList, *p)
	x = x + 1
}

func (productList ProductList) GetIndex(product Product) int {
	for idx, p := range productList {
		if p == product {
			return idx
		}
	}
	return -1
}

func (productList ProductList) Includes(product Product) bool {
	return productList.GetIndex(product) >= 0
}

func (productList ProductList) All(criteria func(Product) bool) bool {
	for _, x := range productList {
		if !criteria(x) {
			return false
		}
	}
	return true
}

func (productList *ProductList) Any(criteria func(Product) bool) bool {
	for _, x := range *productList {
		if criteria(x) {
			return true
		}
	}
	return false
}
