.PHONY: postgres adminer migrate

compose:
	docker-compose -f docker-compose.yml up

migrate:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

migrate1:
	migrate -database postgres://postgres:secret@localhost/postgres?sslmode=disable \
			-path migrations up

migrate-down:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable down