package projeto

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Conexao() {
	// Conectar ao banco de dados PostgreSQL
	var err error
	dsn := "user=estagiario password=projeto_de_estagio dbname=bd_aventura sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Configuração do pool de conexões
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxOpenConns(10)                 // Número máximo de conexões abertas
	sqlDB.SetMaxIdleConns(5)                  // Número máximo de conexões inativas no pool
	sqlDB.SetConnMaxLifetime(time.Minute * 5) // Tempo máximo de vida de uma conexão no pool

	// Cria o banco de dados e as tabelas se não existirem
	if err := criarBancoDeDados(); err != nil {
		log.Fatal(err)
	}

	// Aqui você pode adicionar sua lógica de servidor HTTP
}

func criarBancoDeDados() error {
	// Automigrar o esquema de banco de dados
	err := db.AutoMigrate(&Missao{}, &Aventureiro{}, &Aventura{})
	if err != nil {
		return err
	}

	// Abre o arquivo SQL com as queries para criar as tabelas
	sqlFile, err := os.Open("bd_aventura.sql")
	if err != nil {
		return err
	}
	defer sqlFile.Close()

	// Executa as queries para criar as tabelas
	scanner := bufio.NewScanner(sqlFile)
	for scanner.Scan() {
		query := scanner.Text()
		query = strings.TrimSpace(query)
		if query != "" {
			if err := db.Exec(query).Error; err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
