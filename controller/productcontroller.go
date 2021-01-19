package controller

import (
	"github.com/productweb/models"
	"gorm.io/gorm"
)

type ControllerContextDB struct {
	db *gorm.DB
}

//BuscarProdutos retorna um produto de Id Informado
func (c *ControllerContextDB) BuscarProdutos(id int32) models.Product {
	db := c.db.GetDatabase()
	product := Product{}

	db.First(&product, id)
}
