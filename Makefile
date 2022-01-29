dcup: ## start in development mode
	ENV=development docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

build: ## build the image and start in development mode
	ENV=development docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build

dcdown:
	docker-compose down

initdb:
	docker exec -it datalearn_db_1 createdb --username=root --owner=root datalearn

dropdb:
	docker exec -it datalearn_db_1 dropdb datalearn

migrateup:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datalearn?sslmode=disable" up

migratedown:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datalearn?sslmode=disable" down

.PHONY: dcup build dcdown initdb dropdb migrateup migratedown
