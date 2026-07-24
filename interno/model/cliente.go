package model

// Dados do cliente
type Cliente struct {
	Idade  int
	Name   string
	Email  string
	Cidade string
}

func NewCliente(idade int, name, email, cidade string) *Cliente {
	return &Cliente{
		Idade:  idade,
		Name:   name,
		Email:  email,
		Cidade: cidade,
	}
}