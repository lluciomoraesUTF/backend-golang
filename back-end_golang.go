package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	//Inicia a conexão com o banco de dados
	db, err := sql.Open("postgres", "user=estagiario password=projeto_de_estagio dbname=bd_aventura sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Cria o banco de dados e as tabelas se não existirem
	if err := createDatabase(); err != nil {
		log.Fatal(err)
	}

	// Inicia o servidor na porta 8080
	log.Println("Servidor iniciado na porta 8080")
	// Aqui você pode adicionar sua lógica de servidor HTTP
}

func createDatabase() error {
	//Cria o banco de dados caso ele não exista
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS bd_aventura")
	if err != nil {
		return err
	}

	//Seleciona o banco de dados.
	_, err = db.Exec("bd_aventura")
	if err != nil {
		return err
	}

	// Abre o banco de dados criado.
	sqlFile, err := os.Open("bd_aventura.sql")
	if err != nil {
		return err
	}
	defer sqlFile.Close()

	// Executas as Querys criadas no banco
	scanner := bufio.NewScanner(sqlFile)
	for scanner.Scan() {
		query := scanner.Text()
		query = strings.TrimSpace(query)
		if query != "" {
			_, err := db.Exec(query)
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
