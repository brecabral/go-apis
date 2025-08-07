package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/brecabral/go-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	name := "produto"
	price := 10
	product, err := entity.NewProduct(name, price)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	for i := 1; i <= 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("produto %d", i), rand.Int())
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "produto 1", products[0].Name)
	assert.Equal(t, "produto 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "produto 11", products[0].Name)
	assert.Equal(t, "produto 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "produto 21", products[0].Name)
	assert.Equal(t, "produto 25", products[4].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	name := "produto"
	price := 10
	product, _ := entity.NewProduct(name, price)
	db.Create(product)

	productDB := NewProduct(db)
	productFound, err := productDB.FindByID(product.ID.String())

	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.True(t, product.CreatedAt.Equal(productFound.CreatedAt))
}
