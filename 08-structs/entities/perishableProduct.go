package entities

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProject(id int, name string, cost float32, units int, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{Product{id, name, cost, units, category, true, "dummy perishable product"}, expiry}
}
