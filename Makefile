dev:
	ENV=development docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

api:
	docker run --name data-api -d -p 3100:3100 data-api:latest

postgres:
	docker run --name data-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=spartak1 -d postgres:13.1-alpine

# Container names in commands postgres, createdb and dropdb are not right
createdb:
	docker exec -it data-db createdb --username=root --owner=root datacademy

dropdb:
	docker exec -it data-db dropdb datacademy

migrateup:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datacademy?sslmode=disable" up

migratedown:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datacademy?sslmode=disable" down

.PHONY: dev api postgres createdb dropdb migrateup migratedown