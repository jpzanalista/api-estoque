# api-estoque

API REST de **controle de estoque de produtos**, desenvolvida em **Go** com o
framework web **Gin**.

O projeto foi construído a partir do tutorial oficial
[Desenvolvendo uma API RESTful com Go e Gin](https://go.dev/doc/tutorial/web-service-gin),
adaptado para o tema de gestão de estoque. Cada etapa foi registrada em um
commit, praticando o versionamento com Git.

## Sobre a API

A API gerencia uma coleção de produtos — cada um com id, nome, preço, quantidade
e categoria. Os dados são mantidos em memória (sem banco de dados).

### Endpoints

| Método | Caminho         | Descrição                    |
|--------|-----------------|------------------------------|
| GET    | `/produtos`     | Lista todos os produtos      |
| GET    | `/produtos/:id` | Busca um produto pelo seu id |
| POST   | `/produtos`     | Cadastra um novo produto     |

## Tecnologias

- [Go](https://go.dev/) 1.26
- [Gin](https://github.com/gin-gonic/gin) — framework web
- Git e GitHub para controle de versão

## Como executar

Pré-requisito: ter o [Go instalado](https://go.dev/doc/install).

```bash
git clone https://github.com/jpzanalista/api-estoque.git
cd api-estoque
go run .
```

O servidor sobe em `localhost:8080`. Em outro terminal, teste os endpoints:

```bash
# Listar todos os produtos
curl http://localhost:8080/produtos

# Buscar um produto pelo id
curl http://localhost:8080/produtos/1

# Cadastrar um novo produto
curl http://localhost:8080/produtos \
  --header "Content-Type: application/json" \
  --request POST \
  --data '{"id":"4","nome":"Webcam HD","preco":159.90,"quantidade":12,"categoria":"Periféricos"}'
```

## O que o projeto demonstra

- Construção de uma API REST com o framework Gin;
- Roteamento de requisições por método HTTP e caminho;
- Conversão de structs Go para JSON e vice-versa;
- Leitura de parâmetros de URL e do corpo de requisições;
- Uso de códigos de status HTTP (200, 201, 404);
- Versionamento incremental com Git, endpoint a endpoint.