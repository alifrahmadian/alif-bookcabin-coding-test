APP_NAME=alif_bookcabin_coding_test
MAIN=./cmd

MIGRATION_PATH = db/migrations

MIGRATE = migrate -path $(MIGRATION_PATH)

.PHONY: run migrate-create migrate-up migrate-down migrate-force migrate-version web web-install web-build web-clean

run:
	go run $(MAIN)

migrate-create:
ifndef NAME
	$(error NAME is undefined. Usage: make migrate-create NAME=create_table)
endif
	migrate create -ext sql -dir "$(MIGRATION_PATH)" -seq $(NAME)

web:
	cd web/alif-bookcabin-coding-test && npm run dev

web-install:
	cd web/alif-bookcabin-coding-test && npm install

web-build:
	cd web/alif-bookcabin-coding-test && npm run build

web-clean:
	cd web/alif-bookcabin-coding-test && rm -rf node_modules
