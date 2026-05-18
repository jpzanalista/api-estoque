package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

// getProdutos responde com a lista de todos os produtos em JSON.
func getProdutos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, produtos)
}

// getProdutoPorID localiza o produto cujo id corresponde ao
// parâmetro enviado pelo cliente e o retorna como resposta.
func getProdutoPorID(c *gin.Context) {
	id := c.Param("id")

	// Percorre a lista de produtos procurando um cujo id
	// seja igual ao parâmetro recebido.
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

	// BindJSON converte o JSON do corpo da requisição em um produto.
	if err := c.BindJSON(&novoProduto); err != nil {
		return
	}

	// Adiciona o novo produto à lista.
	produtos = append(produtos, novoProduto)
	c.IndentedJSON(http.StatusCreated, novoProduto)
}

func main() {
	router := gin.Default()
	router.GET("/produtos", getProdutos)
	router.GET("/produtos/:id", getProdutoPorID)
	router.POST("/produtos", postProdutos)

	router.Run("localhost:8080")
}
