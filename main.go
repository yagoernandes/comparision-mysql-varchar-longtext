package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/exp/rand"
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
	db, err := makeConnection()
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

func makeConnection() (*sql.DB, error) {
	// Configuração de conexão
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASSWORD, HOST, PORT, DBNAME)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
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

func insertRandomData(db *sql.DB, tableName string) {
	rand.Seed(uint64(time.Now().UnixNano()))
	tags := generateRandomTags()

	var query string
	switch tableName {
	case "test_varchar":
		query = fmt.Sprintf("INSERT INTO %s (tags) VALUES (?)", tableName)
	case "test_longtext":
		query = fmt.Sprintf("INSERT INTO %s (tags) VALUES (?)", tableName)
	case "test_json":
		query = fmt.Sprintf("INSERT INTO %s (tags) VALUES (JSON_ARRAY(?))", tableName)
	default:
		log.Printf("Tabela %s não suportada para inserção", tableName)
		return
	}

	_, err := db.Exec(query, tags)
	if err != nil {
		log.Printf("Erro ao inserir dados na tabela %s: %v", tableName, err)
		return
	}

	log.Printf("Dados inseridos com sucesso na tabela %s: %v", tableName, tags)
}

// Função que gera dados aleatórios
func generateRandomTags() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	tags := make([]rune, 10)
	for i := range tags {
		tags[i] = letters[rand.Intn(len(letters))]
	}
	return string(tags)
}
