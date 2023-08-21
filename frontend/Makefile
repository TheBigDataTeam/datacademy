build:
	docker build -t data-app:1.0.0 .

app:
	docker run --name data-app -d -p 3000:3000 data-app:1.0.0

.PHONY: build app
