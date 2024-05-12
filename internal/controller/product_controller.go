package controller

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/xprazak2/ventrata/internal/service"
	"github.com/xprazak2/ventrata/internal/view"
)

func GetProduct(svc *service.ProductService) func(w http.ResponseWriter, t *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isPriced := IsPriced(r.Context())
		productId := chi.URLParam(r, "id")

		product := svc.GetProduct(productId)
		if product == nil {
			ResponseError(w, fmt.Sprintf("Product with id '%s' not found", productId), http.StatusNotFound)
			return
		}

		ResponseJSON(w, view.ProductView(*product, isPriced), http.StatusOK)
	}
}

func GetProducts(svc *service.ProductService) func(w http.ResponseWriter, t *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isPriced := IsPriced(r.Context())
		products := svc.GetProducts()
		ResponseJSON(w, view.ProductsView(products, isPriced), http.StatusOK)
	}
}
