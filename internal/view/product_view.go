package view

import "github.com/xprazak2/ventrata/internal/model"

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Capacity uint64 `json:"capacity"`
}

type PricedProduct struct {
	Product
	Price    uint64 `json:"price"`
	Currency string `json:"currency"`
}

func ProductView(product model.Product, isPriced bool) interface{} {
	prod := Product{
		Id:       product.Id,
		Name:     product.Name,
		Capacity: product.Capacity,
	}

	if !isPriced {
		return prod
	}

	return PricedProduct{
		Product:  prod,
		Price:    product.Price,
		Currency: product.Currency,
	}
}

func ProductsView(products []model.Product, isPriced bool) []interface{} {
	res := make([]interface{}, len(products))

	for idx, product := range products {
		res[idx] = ProductView(product, isPriced)
	}

	return res
}
