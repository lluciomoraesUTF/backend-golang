package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func teste_Listar_missoes(t *testing.T) {
	// Configuração do servidor de teste
	r := setupRouter()

	// Simula uma solicitação HTTP GET para /missions
	req, _ := http.NewRequest("GET", "/missions", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verifica o código de status da resposta
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCriarMissao(t *testing.T) {
	// Configuração do servidor de teste
	r := setupRouter()

	// Dados da missão a ser criada
	missao := Missao{Nome_Missao: "Missão de Teste", Descricao: "Descrição da Missão de Teste", Dificuldade: "S"}
	payload, _ := json.Marshal(missao)

	// Simula uma solicitação HTTP POST para /missions
	req, _ := http.NewRequest("POST", "/missions", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verifica o código de status da resposta
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestAtualizarMissao(t *testing.T) {
	// Configuração do servidor de teste
	r := setupRouter()

	// Dados da missão a ser atualizada
	missao := Missao{Nome_Missao: "Missão de Teste Atualizada", Descricao: "Descrição da Missão de Teste Atualizada", Dificuldade: "A"}
	payload, _ := json.Marshal(missao)

	// Simula uma solicitação HTTP PUT para /missions/:id_missao
	req, _ := http.NewRequest("PUT", "/missions/1", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verifica o código de status da resposta
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletarMissao(t *testing.T) {
	// Configuração do servidor de teste
	r := setupRouter()

	// Simula uma solicitação HTTP DELETE para /missions/:id_missao
	req, _ := http.NewRequest("DELETE", "/missions/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verifica o código de status da resposta
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListarAventuras(t *testing.T) {
	// Configuração do servidor de teste
	r := setupRouter()

	// Simula uma solicitação HTTP GET para /adventures
	req, _ := http.NewRequest("GET", "/adventures", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verifica o código de status da resposta
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCriarAventura(t *testing.T) {
	// Configuração do servidor de teste
	r := setupRouter()

	// Dados da aventura a ser criada
	aventura := Aventura{id_missao: 1, id_aventureiro: 1}
	payload, _ := json.Marshal(aventura)

	// Simula uma solicitação HTTP POST para /adventures
	req, _ := http.NewRequest("POST", "/adventures", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verifica o código de status da resposta
	assert.Equal(t, http.StatusCreated, w.Code)
}
func setupRouter() *gin.Engine {
	// Configuração do servidor de teste
	r := gin.Default()

	// Rotas de teste
	r.GET("/missions", listar_MIssoes)
	r.POST("/missions", criar_Missao)
	r.PUT("/missions/:id_missao", atualizar_Missao)
	r.DELETE("/missions/:id_missao", deletar_Missao)
	r.GET("/adventures", listar_Aventuras)
	r.POST("/adventures", criar_Aventura)

	return r
}
