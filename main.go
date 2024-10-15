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
	start := time.Now()
	queryDB(db, "test_varchar", 10)
	fmt.Printf("Tempo de execução test_varchar: %v\n", time.Since(start))

	// Função para consultar a tabela test_longtext
	start = time.Now()
	queryDB(db, "test_longtext", 10)
	fmt.Printf("Tempo de execução test_longtext: %v\n", time.Since(start))

	// Função para consultar a tabela test_json
	start = time.Now()
	queryDB(db, "test_json", 10)
	fmt.Printf("Tempo de execução test_json: %v\n", time.Since(start))
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
func queryDB(db *sql.DB, tableName string, limit int) {
	query := fmt.Sprintf("SELECT id, tags FROM %s ORDER BY id DESC LIMIT %d", tableName, limit)
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
	randomFloat := rand.Float64() * 1000000
	tags = []rune(fmt.Sprintf("{\"%s\": %v}", string(tags), randomFloat))
	return string(tags)
}

// mysqlslap --user=root --password=root --host=localhost \
//           --concurrency=100 --iterations=20 --number-of-queries=1000 \
//           --query="INSERT INTO accumulator (mod_accum, accum_id, accum_owner_id, accum_window, sum_value, version, created_at, updated_at, db ,count_value, min_seq_id, max_seq_id, last_batch_oldest_published_time, last_batch_average_published_time, last_batch_oldest_start_processing_time, last_batch_average_start_processing_time, tags_string) VALUES (88, '1', '807897171', 1711335600, '3000.00000000', 1, '2024-03-25 04:14:26.328457641', '2024-03-25 04:14:26.328457711', 'accumalph1', 1, 473426081, 473426081, '2024-03-25 04:14:24', '2024-03-25 04:14:24', '2024-03-25 04:14:24', '2024-03-25 04:14:24', '{\"approved\":3000.00000000}') ON DUPLICATE KEY UPDATE  sum_value = (VALUES(sum_value) + sum_value), count_value = (VALUES(count_value) + count_value), version = (version + 1), updated_at = VALUES(updated_at), max_seq_id = VALUES(max_seq_id), last_batch_oldest_published_time = VALUES(last_batch_oldest_published_time), last_batch_average_published_time = VALUES(last_batch_average_published_time), last_batch_oldest_start_processing_time = VALUES(last_batch_oldest_start_processing_time),

func insertRandomDataBatch(db *sql.DB, tableName string, batchSize int) {
	rand.Seed(uint64(time.Now().UnixNano()))

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

	// Iniciar uma transação
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Erro ao iniciar transação: %v", err)
		return
	}

	// Preparar a declaração de inserção
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Printf("Erro ao preparar declaração de inserção: %v", err)
		return
	}

	// Executar a declaração de inserção em lote
	for i := 0; i < batchSize; i++ {
		tags := generateRandomTags()
		_, err = stmt.Exec(tags)
		if err != nil {
			log.Printf("Erro ao executar declaração de inserção em lote: %v", err)
			return
		}
	}

	// Commit da transação
	err = tx.Commit()
	if err != nil {
		log.Printf("Erro ao fazer commit da transação: %v", err)
		return
	}

	// Fechar a declaração de inserção
	err = stmt.Close()
	if err != nil {
		log.Printf("Erro ao fechar declaração de inserção: %v", err)
		return
	}
}
