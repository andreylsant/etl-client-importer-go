package csv

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/andreylsant/etldecliente/interno/etl"
	"github.com/andreylsant/etldecliente/interno/model"
)

type ReaderEtl struct{
	transforma etl.Transformar
}

// Essa func inicia uma leitura csv
// Tipo alterado de [][]string para []model.Cliente
func (retl *ReaderEtl) Reader(r io.Reader) ([]model.Cliente, error) {
	fmt.Println("[Iniciando func Reader] Iniciando leitura do file!" )
	//Esta linha inicializa o leitor de arquivos CSV da biblioteca padrão do Go
	reader := csv.NewReader(r)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	// Agora o tipo coincide perfeitamente
	
	cliente, err := retl.transforma.TransformarClientes(records)
	if err != nil {
		return nil, err
	}

	return cliente, nil
}
