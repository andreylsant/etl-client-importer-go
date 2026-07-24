# ETL Client Importer

Pipeline ETL desenvolvido em **Go (Golang)** para importar dados de clientes a partir de arquivos CSV, validar os registros, transformá-los em entidades de domínio e persisti-los em um banco de dados PostgreSQL.

## 📖 Sobre o projeto

Este projeto simula um cenário real de importação de dados, seguindo as etapas de um processo ETL (**Extract, Transform, Load**).

O objetivo é demonstrar boas práticas de desenvolvimento backend com Go, como separação de responsabilidades, organização em camadas, validação de dados e preparação para persistência em banco de dados.

## 🚀 Tecnologias

- Go (Golang)
- PostgreSQL
- Docker
- Docker Compose
- CSV
- SQL
- Git

## 🏗️ Arquitetura

```
CSV
 │
 ▼
Reader
 │
 ▼
Validator
 │
 ▼
Transformer (ETL)
 │
 ▼
Repository
 │
 ▼
PostgreSQL
```

## 📂 Estrutura do projeto

```
.
├── cmd/
│   └── main.go
├── data/
│   └── cliente.csv
├── internal/
│   ├── csv/
│   ├── etl/
│   ├── model/
│   ├── repository/
│   └── validator/
├── go.mod
└── README.md
```

## ⚙️ Fluxo da aplicação

1. A aplicação lê um arquivo CSV.
2. Cada registro é validado.
3. Os dados são transformados em entidades (`Cliente`).
4. Os registros válidos são enviados para a camada de persistência.
5. Os dados são armazenados no PostgreSQL.

## ✅ Funcionalidades

- Leitura de arquivos CSV
- Processamento de registros
- Validação dos dados
- Transformação para entidades de domínio
- Estrutura em camadas
- Tratamento de erros
- Preparado para integração com PostgreSQL

## 📌 Exemplo do arquivo CSV

```csv
idade,nome,email,cidade
30,João,joao@email.com,Rio de Janeiro
25,Maria,maria@email.com,São Paulo
40,Carlos,carlos@email.com,Belo Horizonte
```

## ▶️ Executando o projeto

Clone o repositório:

```bash
git clone https://github.com/andreylsant/etl-client-importer-go.git
```

Entre na pasta:

```bash
cd etl-client-importer-go
```

Execute:

```bash
go run ./cmd
```

## 📈 Próximas melhorias

- [ ] Persistência em PostgreSQL
- [ ] Docker Compose
- [ ] Testes unitários
- [ ] Logs estruturados
- [ ] Relatório de importação
- [ ] Configuração por variáveis de ambiente
- [ ] Validação avançada de e-mail
- [ ] Importação em lote (Batch Insert)

## 📚 Conceitos aplicados

- ETL (Extract, Transform, Load)
- Clean Code
- Responsabilidade Única (SRP)
- Separação de Camadas
- Tratamento de Erros
- Manipulação de Arquivos
- Leitura de CSV
- Conversão de Tipos
- Organização de Projetos em Go

## 👨‍💻 Autor

**Andrey Lopes**

Desenvolvedor Backend em formação, com foco em Go (Golang), APIs REST, PostgreSQL, Docker e Arquitetura de Software.

- GitHub: https://github.com/andreylsant
- LinkedIn: https://www.linkedin.com/in/andrey-lopes-20b3161a7