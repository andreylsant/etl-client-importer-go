package etl

import (
	"fmt"

	"github.com/andreylsant/etldecliente/interno/model"
)

type Transformar struct {
	Cliente []model.Cliente
}

func (t *Transformar) TransformarClientes(registros [][]string) ([]model.Cliente, error) {
	fmt.Println("[Iniciando func TransformaCliente] Iniciando registro")
	for _, registro := range registros {
		cliente := model.Cliente{}

		cliente.Name = registro[0]
		cliente.Email = registro[1]
		cliente.Idade = registro[2]
		cliente.Cidade = registro[3]

			// Antes de retorna o cliente seria melhor validar se os valores passado não são nulos
		err := ValidarRegistro(registro)
		if err != nil {
			return nil, err
		}

		// SALVA o cliente atual na lista
		t.Cliente = append(t.Cliente, cliente)
	}

	return t.Cliente, nil
}
