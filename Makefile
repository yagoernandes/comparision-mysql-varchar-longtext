# Variáveis de configuração
DB_USER=user
DB_PASSWORD=root
DB_HOST=127.0.0.1
DB_PORT=3360
DB_NAME=dbtest

# Caminho das migrations
MIGRATIONS_PATH=./db/migrations

# Alvo para executar as migrations
migrate:
	@for file in $(MIGRATIONS_PATH)/*.sql; do \
		echo "Executando migration $$file"; \
		mysql -h $(DB_HOST) -P $(DB_PORT) -u $(DB_USER) -p$(DB_PASSWORD) $(DB_NAME) < $$file; \
	done

run:
	echo "Executando aplicação"
	go run ./...

.PHONY: migrate run
