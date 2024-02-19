package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func conectar_Banco_Dados() {
	// Conectar ao banco de dados PostgreSQL
	var err error
	db, err = gorm.Open(postgres.Open("user=estagiario password=projeto_de_estagio dbname=bd_aventura sslmode=disable"), &gorm.Config{})
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
}
