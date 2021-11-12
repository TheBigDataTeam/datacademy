dev:
	ENV=development docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

migrateup:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datacademy?sslmode=disable" up

migratedown:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datacademy?sslmode=disable" down

.PHONY: dev migrateup migratedown