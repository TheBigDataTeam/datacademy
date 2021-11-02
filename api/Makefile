.PHONY: build
build:
	docker build -t datalearn-api:latest .

.PHONY: tag
tag:
	docker tag datalearn-api:latest 067789811623.dkr.ecr.us-east-2.amazonaws.com/datalearn-api:latest

.PHONY: push
push:
	docker push 067789811623.dkr.ecr.us-east-2.amazonaws.com/datalearn-api:latest
