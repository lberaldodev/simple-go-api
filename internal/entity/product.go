package entity

import (
	"errors"
	"go_api/pkg/entity"
	"time"
)

var (
	errorNameRequired  = errors.New("name is required")
	errorPriceRequired = errors.New("price is required")
	errorInvalidPrice  = errors.New("price invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		Name:      name,
		Price:     price,
		ID:        entity.NewId(),
		CreatedAt: time.Now(),
	}

	err := product.ValidateProduct()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) ValidateProduct() error {
	if p.Name == "" {
		return errorNameRequired
	}

	if p.Price == 0 {
		return errorPriceRequired
	}

	if p.Price < 0 {
		return errorInvalidPrice
	}
	return nil
}
