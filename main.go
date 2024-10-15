package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USER     = "root"
	PASSWORD = "root"
	HOST     = "localhost"
	PORT     = "3360"
	DBNAME   = "dbtest"
)

func main() {
	// Configuração de conexão
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASSWORD, HOST, PORT, DBNAME)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Testa a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Função para consultar a tabela test_varchar
	queryAndPrint(db, "test_varchar")

	// Função para consultar a tabela test_longtext
	queryAndPrint(db, "test_longtext")

	// Função para consultar a tabela test_json
	queryAndPrint(db, "test_json")
}

// Função para realizar a query e imprimir os resultados
func queryAndPrint(db *sql.DB, tableName string) {
	query := fmt.Sprintf("SELECT id, tags FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Erro ao consultar a tabela %s: %v", tableName, err)
		return
	}
	defer rows.Close()

	fmt.Printf("Resultados da tabela %s:\n", tableName)
	for rows.Next() {
		var id int
		var tags string
		if err := rows.Scan(&id, &tags); err != nil {
			log.Println("Erro ao escanear resultado:", err)
			return
		}
		fmt.Printf("ID: %d, Tags: %s\n", id, tags)
	}

	if err := rows.Err(); err != nil {
		log.Println("Erro ao iterar pelos resultados:", err)
	}
}
