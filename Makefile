dev:
	ENV=development docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

devb:
	ENV=development docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build

stop:
	docker-compose down

initdb:
	docker exec -it datalearn_db_1 createdb --username=root --owner=root datalearn

dropdb:
	docker exec -it datalearn_db_1 dropdb datalearn

migrateup:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datalearn?sslmode=disable" up

migratedown:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datalearn?sslmode=disable" down

.PHONY: dev devb stop initdb dropdb migrateup migratedown