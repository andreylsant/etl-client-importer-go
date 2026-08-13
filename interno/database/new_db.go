package database

import (
	"gorm.io/drive/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=meu_usuario password=minha_senha dbname=meu_banco port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error ao levantar o banco de dados")
	}

	return db
}
