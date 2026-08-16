package etl

import (
	"errors"
	"fmt"
)

func ValidarRegistro(registro []string) error {
	fmt.Println("[Validar Registro]")

	if len(registro) != 4 {
		return errors.New("[Error ao Validar registro] Quantidade de colunas incorreta!")
	}
	//Essa parte é para validar se os campos retornados estão vazios
	
	return nil
}
