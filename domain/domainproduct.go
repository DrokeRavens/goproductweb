package domain

import (
	"github.com/productweb/domain/context"
	"github.com/productweb/models"
)

type productDomain struct{}

type productDomainService interface {
	Buscar(int32) *models.Product
}

var (
	ProductDomainService productDomainService
	dbContext            = context.GetDatabase()
)

func init() {
	ProductDomainService = &productDomain{}
}

func (p *productDomain) Buscar(id int32) *models.Product {
	rtnProduct := models.Product{}
	dbContext.Find(&rtnProduct, id)
	return &rtnProduct
}
