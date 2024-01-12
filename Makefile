IMAGE_NAME = golang-portfolio
VERSION = latest
PORT = 8080
CONTAINER_NAME = golang-portfolio

build-image:
	docker build -t $(IMAGE_NAME):$(VERSION) .

run-container:
	docker run -d -p $(PORT):8080 --name $(CONTAINER_NAME) $(IMAGE_NAME):$(VERSION)

run-container-watch:
	docker run -d --rm -p $(PORT):8080 -v "$(PWD):/app" --name $(CONTAINER_NAME) $(IMAGE_NAME):$(VERSION)

.PHONY: build-image run-container run-container-watch