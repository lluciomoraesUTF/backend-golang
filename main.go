package main

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Definindo a estrutura do modelo para Missão
type Missao struct {
	ID_Missao   uint   `gorm:"primaryKey"`
	Nome_Missao string `json:"nome_missao"`
	Descricao   string `json:"descricao"`
	Dificuldade string `json:"dificuldade"`
}

// Definindo a estrutura do modelo para Aventureiro
type Aventureiro struct {
	ID_Aventureiro uint   `gorm:"primaryKey"`
	Nome_Avent     string `json:"nome_avent"`
	Rank_Avent     string `json:"rank_avent"`
}

// Definindo a estrutura do modelo para Aventura
type Aventura struct {
	ID_Avent       uint `gorm:"primaryKey"`
	ID_Missao      uint `json:"id_missao"`
	ID_Aventureiro uint `json:"id_aventureiro"`
}

var (
	usuario = "estagiario"
	senha   = "projeto_de_estagio"
)

// Middleware de autenticação básica
func Autenticacao() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o header "Authorization" da requisição
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// Se não houver header "Authorization", retorna status de erro
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais não fornecidas"})
			c.Abort()
			return
		}

		// Verifica se o header "Authorization" está no formato "Basic <credenciais_base64>"
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Basic" {
			// Se o formato estiver incorreto, retorna status de erro
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de autenticação inválido"})
			c.Abort()
			return
		}

		// Decodifica as credenciais
		decodedBytes, err := base64.StdEncoding.DecodeString(splitToken[1])
		if err != nil {
			// Se houver um erro ao decodificar, retorna status de erro
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Erro ao decodificar as credenciais"})
			c.Abort()
			return
		}
		credenciais := strings.Split(string(decodedBytes), ":")
		if len(credenciais) != 2 || credenciais[0] != usuario || credenciais[1] != senha {
			// Se as credenciais forem inválidas, retorna status de erro
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
			c.Abort()
			return
		}

		// Se as credenciais estiverem corretas, permite o acesso ao próximo middleware
		c.Next()
	}
}

func main() {
	// Conectar ao banco de dados PostgreSQL
	Conexao()

	// Criando uma instância do Gin
	r := gin.Default()
	//APlica Middleware para autenticar todas as rotas
	r.Use(Autenticacao())
	// Rotas da API
	// Para Missões
	r.GET("/missions", listar_Missao)
	r.GET("/missions/:id_missao", ver_Missao)
	r.POST("/missions", criar_Missao)
	r.PUT("/missions/:id_missao", atualizar_Missao)
	r.DELETE("/missions/:id_missao", deletar_Missao)

	// Para Aventureiros
	r.GET("/adventurers", listar_Aventureiro)
	r.GET("/adventurers/:id_aventureiro", ver_Aventureiro)
	r.POST("/adventurers", criar_Aventureiro)
	r.PUT("/adventurers/:id_aventureiro", atualizar_Aventureiro)
	r.DELETE("/adventurers/:id_aventureiro", deletar_Aventureiro)
	//Aventuras
	//Listar todas as aventuras
	r.GET("/adventures", listar_Aventura)
	// Criar uma nova aventura
	r.POST("/adventures", criar_Aventura)

	// Executar o servidor na porta 8080
	r.Run(":8080")
}

// Funções CRUD para Missões

func listar_Missao(c *gin.Context) {
	var missions []Missao
	db.Find(&missions)
	c.JSON(http.StatusOK, missions)
}

func ver_Missao(c *gin.Context) {
	var mission Missao
	if err := db.First(&mission, c.Param("id_missao")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Missão não encontrada"})
		return
	}
	c.JSON(http.StatusOK, mission)
}

func criar_Missao(c *gin.Context) {
	var input Missao
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mission := Missao{Nome_Missao: input.Nome_Missao, Descricao: input.Descricao, Dificuldade: input.Dificuldade}
	db.Create(&mission)
	c.JSON(http.StatusCreated, mission)
}

func atualizar_Missao(c *gin.Context) {
	var mission Missao
	if err := db.First(&mission, c.Param("id_missao")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Missão não encontrada"})
		return
	}
	var input Missao
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&mission).Updates(input)
	c.JSON(http.StatusOK, mission)
}

func deletar_Missao(c *gin.Context) {
	var mission Missao
	if err := db.First(&mission, c.Param("id_missao")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Missão não encontrada"})
		return
	}
	db.Delete(&mission)
	c.JSON(http.StatusOK, gin.H{"message": "Missão deletada com sucesso"})
}

// Funções CRUD para Aventureiros

func listar_Aventura(c *gin.Context) {
	var adventures []Aventura
	db.Find(&adventures)
	c.JSON(http.StatusOK, adventures)
}

// Funções CRUD para Aventureiros

func listar_Aventureiro(c *gin.Context) {
	var aventureiros []Aventureiro
	db.Find(&aventureiros)
	c.JSON(http.StatusOK, aventureiros)
}

func ver_Aventureiro(c *gin.Context) {
	var aventureiro Aventureiro
	if err := db.First(&aventureiro, c.Param("id_aventureiro")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aventureiro não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aventureiro)
}

func criar_Aventureiro(c *gin.Context) {
	var input Aventureiro
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	aventureiro := Aventureiro{Nome_Avent: input.Nome_Avent, Rank_Avent: input.Rank_Avent}
	db.Create(&aventureiro)
	c.JSON(http.StatusCreated, aventureiro)
}

func atualizar_Aventureiro(c *gin.Context) {
	var aventureiro Aventureiro
	if err := db.First(&aventureiro, c.Param("id_aventureiro")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aventureiro não encontrado"})
		return
	}
	var input Aventureiro
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&aventureiro).Updates(input)
	c.JSON(http.StatusOK, aventureiro)
}

func deletar_Aventureiro(c *gin.Context) {
	var aventureiro Aventureiro
	if err := db.First(&aventureiro, c.Param("id_aventureiro")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aventureiro não encontrado"})
		return
	}
	db.Delete(&aventureiro)
	c.JSON(http.StatusOK, gin.H{"message": "Aventureiro deletado com sucesso"})
}

// Função para converter o rank em valor numérico
func converterRank(rank string) int {
	switch rank {
	case "S":
		return 4
	case "A":
		return 3
	case "B":
		return 2
	case "C":
		return 1
	default:
		return 0 // Rank desconhecido ou inválido
	}
}

// Função para criar uma aventura
func criar_Aventura(c *gin.Context) {
	var input Aventura
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se o aventureiro tem o rank adequado para a missão
	var missao Missao
	if err := db.First(&missao, input.ID_Missao).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Missão não encontrada"})
		return
	}

	var aventureiro Aventureiro
	if err := db.First(&aventureiro, input.ID_Aventureiro).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aventureiro não encontrado"})
		return
	}

	valorMissao := converterRank(missao.Dificuldade)
	valorAventureiro := converterRank(aventureiro.Rank_Avent)

	if valorAventureiro < valorMissao {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rank do aventureiro é inferior à dificuldade da missão"})
		return
	}

	aventura := Aventura{ID_Missao: input.ID_Missao, ID_Aventureiro: input.ID_Aventureiro}
	db.Create(&aventura)
	c.JSON(http.StatusCreated, aventura)
}
