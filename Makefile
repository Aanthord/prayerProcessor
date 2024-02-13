# Import environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Default image name
DOCKER_IMAGE_NAME ?= prayerprocessor

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test ./...

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):latest .

.PHONY: run-dev-stack
run-dev-stack: docker-build
	# Here you can define the steps to start your local development stack,
	# for example using docker-compose up if you have a docker-compose.yml file.
	# This is a placeholder command.
	docker-compose up -d

.PHONY: all
all: vet test docker-build run-dev-stack

