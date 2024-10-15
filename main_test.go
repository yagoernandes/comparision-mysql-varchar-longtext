package main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
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

func TestQueryAndPrint(t *testing.T) {
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

func BenchmarkQueryAndPrintVarchar(b *testing.B) {
	// Configuração de conexão
	db, err := makeConnection()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Função para consultar a tabela test_varchar
	queryAndPrint(db, "test_varchar")
}

func Benchmark_insertRandomData(b *testing.B) {
	// Configuração de conexão
	db, err := makeConnection()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Função para consultar a tabela test_varchar
	insertRandomData(db, "test_varchar")
}
