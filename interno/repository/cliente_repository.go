package repository

import (
	"github.com/andreylsant/etldecliente/interno/model"
)

type Repository interface {
	SaveCliente(cliente *model.Cliente) error 
}
