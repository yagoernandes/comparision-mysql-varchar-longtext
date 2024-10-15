package main

import (
	"log"
	"testing"
)

// func TestMain(m *testing.M) {
// 	// Configuração de conexão
// 	db, err := makeConnection()
// 	if err != nil {
// 		log.Fatal("Erro ao conectar ao banco de dados:", err)
// 	}
// 	defer db.Close()

// 	// Testa a conexão
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Erro ao conectar ao banco de dados:", err)
// 	}

// 	// Função para consultar a tabela test_varchar
// 	queryDB(db, "test_varchar")

// 	// Função para consultar a tabela test_longtext
// 	queryDB(db, "test_longtext")

// 	// Função para consultar a tabela test_json
// 	queryDB(db, "test_json")
// }

func TestQueryDB(t *testing.T) {
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
	queryDB(db, "test_varchar", 10)

	// Função para consultar a tabela test_longtext
	queryDB(db, "test_longtext", 10)

	// Função para consultar a tabela test_json
	queryDB(db, "test_json", 10)
}

func Benchmark_insertRandomDataBatch_Varchar(b *testing.B) {
	// Configuração de conexão
	db, err := makeConnection()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	insertRandomDataBatch(db, "test_varchar", b.N)
}
func Benchmark_insertRandomDataBatch_Longtext(b *testing.B) {
	// Configuração de conexão
	db, err := makeConnection()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	insertRandomDataBatch(db, "test_longtext", b.N)
}

func BenchmarkQueryDBVarchar(b *testing.B) {
	// Configuração de conexão
	db, err := makeConnection()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Função para consultar a tabela test_varchar
	queryDB(db, "test_varchar", 10)
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
