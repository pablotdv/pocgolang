package routes

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pablotdv/pocgolang/data"
	"github.com/pablotdv/pocgolang/models"
	"github.com/pablotdv/pocgolang/schemas"
)

// @Summary retorna pessoas
// @Success 200 {array} models.Pessoa
// @Router /pessoas [get]
func GetPessoas(c *gin.Context) {
	var pessoas []models.Pessoa
	err := data.Db.Debug().Find(&pessoas).Error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, pessoas)
}

func PostPessoa(c *gin.Context) {
	var pessoaRequest schemas.Pessoa
	err := c.ShouldBindJSON(&pessoaRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pessoa := models.Pessoa{
		Nome:  pessoaRequest.Nome,
		Idade: pessoaRequest.Idade,
	}

	err = data.Db.Debug().Create(&pessoa).Error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, pessoa)
}

func PostSincronizarPessoa(c *gin.Context) {
	var pessoasRequest []schemas.Pessoa
	err := c.ShouldBindJSON(&pessoasRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, pessoaRequest := range pessoasRequest {
		pessoa := models.Pessoa{
			Nome:  pessoaRequest.Nome,
			Idade: pessoaRequest.Idade,
		}

		err = data.Db.Debug().Create(&pessoa).Error
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.Status(http.StatusNoContent)
}

func PostSincronizarPessoa2(c *gin.Context) {
	var pessoasRequest []schemas.Pessoa
	err := c.ShouldBindJSON(&pessoasRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, pessoaRequest := range pessoasRequest {
		go createPessoa2(pessoaRequest)
	}

	c.Status(http.StatusNoContent)
}

func createPessoa2(pessoaRequest schemas.Pessoa) {
	pessoa := models.Pessoa{
		Nome:  pessoaRequest.Nome,
		Idade: pessoaRequest.Idade,
	}

	data.Db.Debug().Create(&pessoa)
}

func PostPesssoaSincronizar3(c *gin.Context) {
	var pessoasRequest []schemas.Pessoa
	err := c.ShouldBindJSON(&pessoasRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var wg sync.WaitGroup

	for _, pessoaRequest := range pessoasRequest {
		go createPessoa3(pessoaRequest, &wg)
	}

	wg.Wait()

	c.Status(http.StatusNoContent)
}

func createPessoa3(pessoaRequest schemas.Pessoa, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	pessoa := models.Pessoa{
		Nome:  pessoaRequest.Nome,
		Idade: pessoaRequest.Idade,
	}

	data.Db.Debug().Create(&pessoa)
}
