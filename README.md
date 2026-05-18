# api-estoque

API REST para controle de estoque de produtos, desenvolvida em Go. Os produtos são armazenados em um banco de dados PostgreSQL, garantindo que os dados persistam entre as execuções.

## Tecnologias

- **Go** — linguagem da aplicação
- **Gin** — framework HTTP e roteamento
- **PostgreSQL** — banco de dados relacional
- **Docker** — execução do banco de dados em contêiner
- **database/sql** com o driver **lib/pq** — acesso ao banco

## Pré-requisitos

- Go instalado
- Docker instalado

## Como executar

### 1. Subir o banco de dados

Crie e inicie um contêiner PostgreSQL:

```
docker run --name estoque-postgres \
  -e POSTGRES_USER=estoque \
  -e POSTGRES_PASSWORD=estoque123 \
  -e POSTGRES_DB=estoque \
  -p 5432:5432 \
  -v estoque-pgdata:/var/lib/postgresql/data \
  -d postgres:17
```

### 2. Criar a tabela

Conecte-se ao banco:

```
docker exec -it estoque-postgres psql -U estoque -d estoque
```

E crie a tabela `produtos`:

```sql
CREATE TABLE produtos (
    id         SERIAL PRIMARY KEY,
    nome       VARCHAR(100) NOT NULL,
    preco      DOUBLE PRECISION NOT NULL,
    quantidade INTEGER NOT NULL,
    categoria  VARCHAR(50) NOT NULL
);
```

### 3. Iniciar a API

```
go run .
```

A API ficará disponível em `http://localhost:8080`.

## Endpoints

| Método | Rota            | Descrição                  |
|--------|-----------------|----------------------------|
| GET    | `/produtos`     | Lista todos os produtos    |
| GET    | `/produtos/:id` | Retorna um produto pelo id |
| POST   | `/produtos`     | Cadastra um novo produto   |

No `POST`, o campo `id` é gerado automaticamente pelo banco e não deve ser enviado. Exemplo de corpo da requisição:

```json
{
    "nome": "Webcam HD",
    "preco": 159.90,
    "quantidade": 12,
    "categoria": "Periféricos"
}
```
