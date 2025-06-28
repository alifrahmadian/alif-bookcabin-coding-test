APP_NAME=alif_bookcabin_coding_test
MAIN=./cmd

MIGRATION_PATH = db/migrations

MIGRATE = migrate -path $(MIGRATION_PATH)

.PHONY: run migrate-create migrate-up migrate-down migrate-force migrate-version

run:
	go run $(MAIN)

migrate-create:
ifndef NAME
	$(error NAME is undefined. Usage: make migrate-create NAME=create_table)
endif
	migrate create -ext sql -dir "$(MIGRATION_PATH)" -seq $(NAME)
