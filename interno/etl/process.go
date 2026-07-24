package etl

import (
	"fmt"
	"strconv"

	"github.com/andreylsant/etldecliente/interno/model"
)

type Transformar struct {
	Cliente []model.Cliente
}

func (t *Transformar) TransformarClientes(registros [][]string) ([]model.Cliente, error) {
	fmt.Println("[Iniciando func TransformaCliente] Iniciando registro")
	for _, registro := range registros {
		cliente := model.Cliente{}

		// Antes de retorna o cliente seria melhor validar se os valores passado não são nulos
		err := ValidarRegistro(registro)
		if err != nil {
			return nil, err
		}

		idade, _ := strconv.Atoi(registro[0])

		cliente.Idade = idade
		cliente.Name = registro[1]
		cliente.Email = registro[2]
		cliente.Cidade = registro[3]

		// SALVA o cliente atual na lista
		t.Cliente = append(t.Cliente, cliente)
	}

	return t.Cliente, nil
}
