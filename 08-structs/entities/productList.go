package entities

type ProductList []Product

func (productList *ProductList) AddProduct(p *Product) {
	*productList = append(*productList, *p)
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

//implement the sort (by using the sort.Sort() method )
/*
	initial implementation should sort the products by id
	provide implementations for sorting the products by cost, units, category etc
*/
