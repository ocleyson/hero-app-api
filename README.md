# Hero App API

## Tecnologias usadas

- Golang v1.16
- Gorm v1.21.12
- Mux v1.8.0
- Pgconn v1.9.0
- Godotenv v1.3.0
- Testify v1.7.0

## Instalação

No diretório Home clone o repositório com o comando `git clone https://github.com/ocleyson/hero-app-api`.

Atenção: Você deve ter em seu computador o postgres instalado e criar um database para a api usar.

## Como rodar a API

Crie um arquivo `.env` no root do projeto com as variáveis descritas em `.env.example` e rode a API com o comando `go run .`.

## Como rodar os testes

Entre na pasta `tests` do projeto e crie um arquivo `.env` com as variáveis descritas em `.env.example` e rode os testes com o comando `go test`.