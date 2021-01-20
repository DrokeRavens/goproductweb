package context

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//GetDatabase return the current context
func GetDatabase() *gorm.DB {
	dsn := "sqlserver://guser:guser@localhost:1433/?database=golang"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
