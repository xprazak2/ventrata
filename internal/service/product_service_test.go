package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xprazak2/ventrata/internal/repository/memory"
)

func TestGetProduct(t *testing.T) {
	repo := memory.NewMemoryRepository()

	svc := NewProductService(repo)

	name := "foo"
	var capacity uint64 = 4
	var price uint64 = 200
	currency := "USD"

	prod := repo.CreateEntry(name, capacity, price, currency)

	prod1 := svc.GetProduct(prod.Id)
	assert.Equal(t, name, prod1.Name)
	assert.Equal(t, capacity, prod1.Capacity)
	assert.Equal(t, price, prod1.Price)
	assert.Equal(t, currency, prod1.Currency)

	prod2 := svc.GetProduct("bar")
	assert.Nil(t, prod2)
}

func TestGetProducts(t *testing.T) {
	repo := memory.NewMemoryRepository()

	svc := NewProductService(repo)

	name1 := "foo"
	name2 := "bar"
	var capacity uint64 = 4
	var price uint64 = 200
	currency := "USD"

	repo.CreateEntry(name1, capacity, price, currency)
	repo.CreateEntry(name2, capacity, price, currency)

	prods := svc.GetProducts()
	assert.Len(t, prods, 2)
}
