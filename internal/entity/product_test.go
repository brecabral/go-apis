package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	name := "produto"
	price := 10
	p, err := NewProduct(name, price)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, name, p.Name)
	assert.Equal(t, price, p.Price)
}

func TestProduct_EmptyName(t *testing.T) {
	name := ""
	price := 10
	p, err := NewProduct(name, price)
	assert.Nil(t, p)
	assert.Error(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProduct_EmptyPrice(t *testing.T) {
	name := "produto"
	price := 0
	p, err := NewProduct(name, price)
	assert.Nil(t, p)
	assert.Error(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProduct_NegativePrice(t *testing.T) {
	name := "produto"
	price := -10
	p, err := NewProduct(name, price)
	assert.Nil(t, p)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}
