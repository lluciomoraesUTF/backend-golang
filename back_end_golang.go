package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Definindo a estrutura do modelo para Missão de Aventura
type Missao struct {
	ID          uint `gorm:"primaryKey"`
	NomeMissao  string
	Descricao   string
	Dificuldade string
}

// Definindo a estrutura do modelo para Personagem
type Aventureiro struct {
	ID         uint `gorm:"primaryKey"`
	nome_avent string
	rank_avent string
}

func main() {
	// Conectar ao banco de dados PostgreSQL
	db := ConnectDB()
	// Fechar a conexão com o banco de dados ao final
	defer db.Close()

	// Migrar o schema
	db.AutoMigrate(&Missao{}, &Aventureiro{})

	// Criando uma instância do Gin
	r := gin.Default()

	// Rotas da API
	// Lista todas as missões.
	r.GET("/missions", func(c *gin.Context) {
		var missions []Missao
		db.Find(&missions)
		c.JSON(http.StatusOK, missions)
	})

	// Seleciona uma missão específica
	r.GET("/missions/:id_missao", func(c *gin.Context) {
		var mission Missao
		if err := db.First(&mission, c.Param("id_missao")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Missão não encontrada"})
			return
		}
		c.JSON(http.StatusOK, mission)
	})

	// Cria uma nova missão
	r.POST("/missions", func(c *gin.Context) {
		var input Missao
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mission := Missao{NomeMissao: input.NomeMissao, Descricao: input.Descricao, Dificuldade: input.Dificuldade}
		db.Create(&mission)
		c.JSON(http.StatusCreated, mission)
	})

	//Atualiza uma missão existente
	r.PUT("/missions/:id_missao", func(c *gin.Context) {
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
	})

	// Deleta uma missão existente
	r.DELETE("/missions/:id_missao", func(c *gin.Context) {
		var mission Missao
		if err := db.First(&mission, c.Param("id_missao")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Missão não encontrada"})
			return
		}
		db.Delete(&mission)
		c.JSON(http.StatusOK, gin.H{"message": "Missão deletada com sucesso"})
	})

	// Executar o servidor na porta 8080
	r.Run(":8080")
}

// Função para conectar ao banco de dados
func ConnectDB() *gorm.DB {
	// Conectar ao banco de dados SQLite
	db, err := gorm.Open(sqlite.Open("adventure.db"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}
	return db
}
