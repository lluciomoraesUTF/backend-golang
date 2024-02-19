package main

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Definindo a estrutura do modelo para Missão
type Missao struct {
	id_missao   uint   `gorm:"primaryKey"`
	Nome_Missao string `json:"nome_missao"`
	Descricao   string `json:"descricao"`
	Dificuldade string `json:"dificuldade"`
}

// Definindo a estrutura do modelo para Aventureiro
type Aventureiro struct {
	id_aventureiro uint   `gorm:"primaryKey"`
	Nome_Avent     string `json:"nome_avent"`
	RankAvent      string `json:"rank_avent"`
}

// Definindo a estrutura do modelo para Aventura
type Aventura struct {
	id_avent       uint `gorm:"primaryKey"`
	id_missao      uint `json:"id_missao"`
	id_aventureiro uint `json:"id_aventureiro"`
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
	conectar_Banco_Dados()

	// Migra o esquema de banco de dados
	db.AutoMigrate(&Missao{}, &Aventureiro{}, &Aventura{})

	// Criando uma instância do Gin
	r := gin.Default()
	//APlica Middleware para autenticar todas as rotas
	r.Use(Autenticacao())
	// Rotas da API
	// Para Missões
	r.GET("/missions", listar_MIssoes)
	r.GET("/missions/:id_missao", ver_Missoes)
	r.POST("/missions", criar_Missao)
	r.PUT("/missions/:id_missao", atualizar_Missao)
	r.DELETE("/missions/:id_missao", deletar_Missao)

	// Para Aventureiros
	r.GET("/aventureiros", listar_Aventureiros)
	r.GET("/aventureiros/:id_aventureiro", ver_Aventureiro)
	r.POST("/aventureiros", criar_Aventureiro)
	r.PUT("/aventureiros/:id_aventureiro", atualizar_aventureiro)
	r.DELETE("/aventureiros/:id_aventureiro", deletar_Aventureiro)
	//Aventuras
	//Listar todas as aventuras
	r.GET("/aventuras", listar_Aventuras)
	// Criar uma nova aventura
	r.POST("/aventuras", criar_Aventura)

	// Executar o servidor na porta 8080
	r.Run(":8080")
}

//CRUD para Missões

func listar_MIssoes(c *gin.Context) {
	var missions []Missao
	db.Find(&missions)
	c.JSON(http.StatusOK, missions)
}

func ver_Missoes(c *gin.Context) {
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

func listar_Aventureiros(c *gin.Context) {
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
	aventureiro := Aventureiro{Nome_Avent: input.Nome_Avent, RankAvent: input.RankAvent}
	db.Create(&aventureiro)
	c.JSON(http.StatusCreated, aventureiro)
}

func atualizar_aventureiro(c *gin.Context) {
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
func listar_Aventuras(c *gin.Context) {
	var adventures []Aventura
	db.Find(&adventures)
	c.JSON(http.StatusOK, adventures)
}

// Função para converter o rank em valor numérico
func converter_Rank(rank string) int {
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

// Função para criar uma  aventura
func criar_Aventura(c *gin.Context) {
	var input Aventura
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se o aventureiro tem o rank adequado para a missão
	var missao Missao
	if err := db.First(&missao, input.id_missao).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Missão não encontrada"})
		return
	}

	var aventureiro Aventureiro
	if err := db.First(&aventureiro, input.id_aventureiro).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aventureiro não encontrado"})
		return
	}

	valorMissao := converter_Rank(missao.Dificuldade)
	valorAventureiro := converter_Rank(aventureiro.RankAvent)

	if valorAventureiro < valorMissao {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rank do aventureiro é inferior à dificuldade da missão"})
		return
	}

	aventura := Aventura{id_missao: input.id_missao, id_aventureiro: input.id_aventureiro}
	db.Create(&aventura)
	c.JSON(http.StatusCreated, aventura)
}
