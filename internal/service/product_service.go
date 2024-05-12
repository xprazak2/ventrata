package service

import (
	"github.com/xprazak2/ventrata/internal/model"
	"github.com/xprazak2/ventrata/internal/repository"
)

type ProductService struct {
	repo repository.Repository
}

func NewProductService(repo repository.Repository) *ProductService {
	return &ProductService{repo}
}

func (svc *ProductService) GetProducts() []model.Product {
	return svc.repo.GetProducts()
}

func (svc *ProductService) GetProduct(id string) *model.Product {
	return svc.repo.GetProduct(id)
}
