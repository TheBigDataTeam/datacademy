build:
	docker build -t data-api:1.0.0 .

api:
	docker run --name data-api -d -p 3100:3100 data-api:1.0.0

.PHONY: build api