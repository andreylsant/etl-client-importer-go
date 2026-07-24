package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andreylsant/etldecliente/interno/csv"
)

func main() {
	fmt.Println("Iniciando projeto")
	reader:= csv.ReaderEtl{}

	file, err := os.Open("data/cliente.csv")
	if err != nil {
		log.Fatal("error ao abrir arquivo %v", err)
	}

	readercsv, err := reader.Reader(file)
	if err != nil {
    log.Fatal(err)
	}

	for _, u := range readercsv {
		fmt.Printf("Salvando no banco -> Nome: %s, Email: %s\n", u.Name, u.Email, )
	}
}
