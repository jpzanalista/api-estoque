package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// produto representa os dados de um produto do estoque.
type produto struct {
	ID         string  `json:"id"`
	Nome       string  `json:"nome"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
	Categoria  string  `json:"categoria"`
}

// produtos é a lista que serve como o estoque inicial (dados de exemplo).
var produtos = []produto{
	{ID: "1", Nome: "Teclado Mecânico", Preco: 249.90, Quantidade: 15, Categoria: "Periféricos"},
	{ID: "2", Nome: "Monitor 24 polegadas", Preco: 899.00, Quantidade: 8, Categoria: "Monitores"},
	{ID: "3", Nome: "Mouse sem fio", Preco: 79.90, Quantidade: 40, Categoria: "Periféricos"},
}

// db é a conexão com o banco de dados PostgreSQL.
var db *sql.DB

// getProdutos responde com a lista de todos os produtos em JSON.
func getProdutos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, produtos)
}

// getProdutoPorID localiza o produto cujo id corresponde ao parâmetro.
func getProdutoPorID(c *gin.Context) {
	id := c.Param("id")

	for _, p := range produtos {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "produto não encontrado"})
}

// postProdutos adiciona um produto a partir do JSON recebido no corpo da requisição.
func postProdutos(c *gin.Context) {
	var novoProduto produto

	if err := c.BindJSON(&novoProduto); err != nil {
		return
	}

	produtos = append(produtos, novoProduto)
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
