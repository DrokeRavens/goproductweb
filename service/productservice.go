package service

import (
	"github.com/productweb/domain"
	"github.com/productweb/models"
)

type productService struct {
}

var (
	ProductService productService
)

func (p *productService) Buscar(id int32) *models.Product {
	productDao := domain.ProductDomainService.Buscar(id)
	return productDao
}
