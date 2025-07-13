include .env
export
MIGRATIONS_DIR=./db/migrations


migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "❌ failed create migrations file"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up


migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1


migrate-status:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version


migrate-reset:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" drop -f
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

integration-test:
	go test -v ./tests/main_integration_test.go