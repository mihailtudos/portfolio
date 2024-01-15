PROJECT=portfolio
DB_CONTAINER_NAME=portfolio-postgres-db
IMAGE_NAME = golang-portfolio
VERSION = latest
PORT = 8080
CONTAINER_NAME = golang-portfolio

build-image:
	docker build -t $(IMAGE_NAME):$(VERSION) .

postgres-up:
	docker compose up

postgres-down:
	docker compose down

postgresit:
	docker exec -it $(DB_CONTAINER_NAME) psql $(PORTFOLIO_DB_DSN)

run-container:
	docker run -d -p $(PORT):8080 --name $(CONTAINER_NAME) $(IMAGE_NAME):$(VERSION)

run-container-watch:
	docker run -d --rm -p $(PORT):8080 -v "$(PWD):/app" --name $(CONTAINER_NAME) $(IMAGE_NAME):$(VERSION)

migrateup:
	migrate -path=./migrations -database=$(PORTFOLIO_DB_DSN) up

migratedown:
	migrate -path=./migrations -database=$(PORTFOLIO_DB_DSN) down

migrateversion:
	migrate -path=./migrations -database=$(PORTFOLIO_DB_DSN) version

.PHONY: build-image run-container run-container-watch postgres-up postgresdown migrateup migrateversion