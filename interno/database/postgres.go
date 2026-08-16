package database

import (
	"github.com/andreylsant/etldecliente/interno/model"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func (db *Database) SaveCliente(cliente *model.Cliente) error {
	tx:= db.Db.Create(cliente)
	return tx.Error
}
