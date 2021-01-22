.PHONY: postgres adminer migrate

compose:
	docker-compose -f docker-compose.yml up

migrate:
	migrate -source file://migrations \
			-database postgres://postgres:secret@db:5432/syndicate?sslmode=disable up

migrate-down:
	migrate -source file://migrations \
			-database postgres://postgres:secret@db:5432/postgres?sslmode=disable down