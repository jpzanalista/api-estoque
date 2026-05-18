package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// produto representa os dados de um produto do estoque.
type produto struct {
	ID         int64   `json:"id"`
	Nome       string  `json:"nome"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
	Categoria  string  `json:"categoria"`
}

// db é a conexão com o banco de dados PostgreSQL.
var db *sql.DB

// getProdutos responde com a lista de todos os produtos do banco.
func getProdutos(c *gin.Context) {
	rows, err := db.Query("SELECT id, nome, preco, quantidade, categoria FROM produtos")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	defer rows.Close()

	produtos := []produto{}
	for rows.Next() {
		var p produto
		if err := rows.Scan(&p.ID, &p.Nome, &p.Preco, &p.Quantidade, &p.Categoria); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}
		produtos = append(produtos, p)
	}

	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, produtos)
}

// getProdutoPorID busca no Banco o produto com id informado.
func getProdutoPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "produto não encontrado"})
		return
	}

	var p produto
	row := db.QueryRow("SELECT id, nome, preco, quantidade, categoria FROM produtos WHERE id = $1", id)
	if err := row.Scan(&p.ID, &p.Nome, &p.Preco, &p.Quantidade, &p.Categoria); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "produto não encontrado"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, p)
}

// postProdutos adiciona um produto a partir do JSON recebido.
func postProdutos(c *gin.Context) {
	var novoProduto produto
	if err := c.BindJSON(&novoProduto); err != nil {
		return
	}

	// O INSERT grava o produto; o RETURNING id devolve o id gerado pelo banco.
	err := db.QueryRow(
		"INSERT INTO produtos (nome, preco, quantidade, categoria) VALUES ($1, $2, $3, $4) RETURNING id",
		novoProduto.Nome, novoProduto.Preco, novoProduto.Quantidade, novoProduto.Categoria,
	).Scan(&novoProduto.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, novoProduto)
}

func main() {
	// String de conexão — usa os valores definidos ao criar o contêiner.
	connStr := "host=localhost port=5432 user=estoque password=estoque123 dbname=estoque sslmode=disable"

	// sql.Open prepara o acesso ao banco (ainda não conecta de fato).
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao preparar o banco: ", err)
	}

	// Ping força uma conexão real, confirmando que o banco está acessível.
	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao conectar no banco: ", err)
	}
	log.Println("Conectado ao PostgreSQL com sucesso.")

	router := gin.Default()
	router.GET("/produtos", getProdutos)
	router.GET("/produtos/:id", getProdutoPorID)
	router.POST("/produtos", postProdutos)

	router.Run("localhost:8080")
}
