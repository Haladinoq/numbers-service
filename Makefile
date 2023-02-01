postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=unumbers -e POSTGRES_PASSWORD=cc402fd0 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=unumbers --owner=unumbers numbers

dropdb:
	docker exec -it postgres12 dropdb --username=unumbers numbers

migrationup:
	migrate -database "postgres://unumbers:cc402fd0@localhost:5433/numbers?sslmode=disable" -path "migrations" up

migrationdown:
	migrate -database "postgres://unumbers:cc402fd0@localhost:5433/numbers?sslmode=disable" -path "migrations" down

swagger:
	swag init -g cmd/numbers-service/main.go

server:
	 go run cmd/numbers-service/main.go -config="config/config.yml"


.PHONY: postgres createdb dropdb migrationup migrationdown swagger server