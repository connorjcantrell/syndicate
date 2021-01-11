.PHONY: postgres adminer migrate

compose:
	docker-compose -f docker-compose.yml up

migrate:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost:5432/syndicate?sslmode=disable up

migrate-down:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable down