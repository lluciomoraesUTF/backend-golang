package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Conexao() {
	// Conectar ao banco de dados PostgreSQL
	var err error
	dsn := "user=estagiario password= '123' dbname=bd_aventura sslmode=disable"
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
	// Verifica a existencia das tabelas
	hasMissao := db.Migrator().HasTable(&Missao{})
	hasAventureiro := db.Migrator().HasTable(&Aventureiro{})
	hasAventura := db.Migrator().HasTable(&Aventura{})

	// Se alguma tabela já existir, não cria novas
	if hasMissao || hasAventureiro || hasAventura {
		log.Println("As tabelas já existem. Ignorando a criação.")
		return nil
	}

	// Automigrar o esquema de banco de dados
	err := db.AutoMigrate(&Missao{}, &Aventureiro{}, &Aventura{})
	if err != nil {
		return err
	}

	return nil
}
