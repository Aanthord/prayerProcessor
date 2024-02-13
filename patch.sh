#!/bin/bash

# Define project directory and subdirectories
PROJECT_DIR=~/go/src/github.com/aanthord/prayerprocessor
CMD_DIR=$PROJECT_DIR/cmd
PKG_DIR=$PROJECT_DIR/pkg
MODELS_DIR=$PKG_DIR/models
SERVICES_DIR=$PKG_DIR/services

# Create directories if they don't exist
mkdir -p $CMD_DIR $MODELS_DIR $SERVICES_DIR

# Exit if any command fails
set -e

# Function to create a file if it doesn't exist and write content to it
create_file() {
    FILE=$1
    CONTENT=$2
    if [ ! -f "$FILE" ]; then
        echo "$CONTENT" > $FILE
    fi
}

# .env file content
ENV_CONTENT="KAFKA_BROKERS=localhost:9092\nSERVER_PORT=3000\nJAEGER_ENDPOINT=http://localhost:14268/api/traces\nJAEGER_SERVICE_NAME=prayerProcessorService"
create_file "$PROJECT_DIR/.env" "$ENV_CONTENT"

# Dockerfile content
DOCKERFILE_CONTENT="FROM golang:1.18 as builder\nWORKDIR /app\nCOPY go.mod go.sum ./\nRUN go mod download\nCOPY . .\nRUN go vet ./...\nRUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prayerProcessor .\n\nFROM registry.access.redhat.com/ubi8/ubi-minimal\nRUN microdnf update && microdnf install ca-certificates && microdnf clean all\nWORKDIR /root/\nCOPY --from=builder /app/prayerProcessor .\nCMD [\"./prayerProcessor\"]"
create_file "$PROJECT_DIR/Dockerfile" "$DOCKERFILE_CONTENT"

# docker-compose.yml content
DOCKER_COMPOSE_CONTENT="version: '3'\nservices:\n  zookeeper:\n    image: wurstmeister/zookeeper\n    ports:\n      - \"2181:2181\"\n  kafka:\n    image: wurstmeister/kafka:2.12-2.3.0\n    ports:\n      - \"9092:9092\"\n    environment:\n      KAFKA_ADVERTISED_HOST_NAME: kafka\n      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181\n    volumes:\n      - /var/run/docker.sock:/var/run/docker.sock\n    depends_on:\n      - zookeeper\n  jaeger:\n    image: jaegertracing/all-in-one:latest\n    ports:\n      - \"5775:5775/udp\"\n      - \"6831:6831/udp\"\n      - \"6832:6832/udp\"\n      - \"5778:5778\"\n      - \"16686:16686\"\n      - \"14268:14268\"\n      - \"14250:14250\"\n      - \"9411:9411\"\n    environment:\n      COLLECTOR_ZIPKIN_HTTP_PORT: 9411"
create_file "$PROJECT_DIR/docker-compose.yml" "$DOCKER_COMPOSE_CONTENT"

# Initialize Go module only if it hasn't been initialized
if [ ! -f "$PROJECT_DIR/go.mod" ]; then
    cd "$PROJECT_DIR" && go mod init github.com/aanthord/prayerprocessor
fi

# main.go content
MAIN_GO_CONTENT="package main\nimport \"fmt\"\nfunc main() {\n    fmt.Println(\"Hello, Prayer Processor!\")\n}"
create_file "$CMD_DIR/main.go" "$MAIN_GO_CONTENT"

# Placeholder for additional Go files
# Replace or add actual content as needed for models, services, etc.
MODEL_GO_CONTENT="package models\n\ntype ExampleModel struct {\n    ID int\n}"
create_file "$MODELS_DIR/example_model.go" "$MODEL_GO_CONTENT"

SERVICE_GO_CONTENT="package services\n\nfunc ExampleService() string {\n    return \"This is an example service\"\n}"
create_file "$SERVICES_DIR/example_service.go" "$SERVICE_GO_CONTENT"

# Go mod tidy to clean up dependencies
cd "$PROJECT_DIR" && go mod tidy

# Build Docker container
docker build -t prayerprocessor:latest .

# Run Docker Compose environment
docker-compose up -d

echo "Deployment of Prayer Processor is complete."

