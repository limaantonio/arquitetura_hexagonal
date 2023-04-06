package dto

import "github.com/limaantonio/go-hexagonal/application"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Status string  `json:"status"`
	Price  float64 `json:"price"`
}

func NewProduct(product application.Product) Product {
	return Product{
		ID:     product.ID,
		Name:   product.Name,
		Status: product.Status,
		Price:  product.Price,
	}
}

func (p *Product) Bind(product *application.Product) (application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status
	_, err := product.IsValid()
	if err != nil {
		return application.Product{}, err
	}
	return *product, nil

}
