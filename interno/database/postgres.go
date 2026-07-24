package database

import "github.com/andreylsant/etldecliente/interno/model"

type Database struct{
}

func (db *Database) SaveCliente(cliente *model.Cliente) error{
	return nil
}
