package model

// Dados do cliente
type Cliente struct {
	Name   string
	Email  string
	Idade string
	Cidade string
}

//Acrescentar regras de negocio na nossa entidade.
//Boas proticas 
func NewCliente(idade, name, email, cidade string) *Cliente {
	return &Cliente{
		Name:   name,
		Email:  email,
		Idade:  idade,
		Cidade: cidade,
	}
}