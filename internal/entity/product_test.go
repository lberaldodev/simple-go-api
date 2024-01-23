package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Testando", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, product.Name, "Testando")
	assert.Equal(t, product.Price, 10.0)
}

func TestProductNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, err, errorNameRequired)
}

func TestProductPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Testando", 0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, err, errorPriceRequired)
}

func TestProductPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Testando", -1)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, err, errorInvalidPrice)
}
