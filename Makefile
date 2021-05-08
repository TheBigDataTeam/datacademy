api:
	docker run --name data-api -d -p 3100:3100 data-api:latest

postgres:
	docker run --name data-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=spartak1 -d postgres:13.1-alpine

createdb:
	docker exec -it data-db createdb --username=root --owner=root datalearn

dropdb:
	docker exec -it data-db dropdb datalearn

migrateup:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datalearn?sslmode=disable" up

migratedown:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datalearn?sslmode=disable" down

.PHONY: api postgres createdb dropdb migrateup migratedown