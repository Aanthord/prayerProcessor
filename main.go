package main

import (
    "context"
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
    "log"
    "os"
    "prayerProcessor/api"
    "prayerProcessor/config"
    "prayerProcessor/service"
    _ "prayerProcessor/docs" // Import generated Swagger docs
)

// @title Prayer Processor API
// @description This service processes prayer requests and demonstrates a self-documenting API with Swagger.
// @version 1.0
// @host localhost:3000
// @BasePath /
func initTracer() {
    godotenv.Load() // Load .env file if present

    jaegerEndpoint := os.Getenv("JAEGER_ENDPOINT")
    if jaegerEndpoint == "" {
        jaegerEndpoint = "http://localhost:14268/api/traces" // Default if not set
    }
    serviceName := os.Getenv("JAEGER_SERVICE_NAME")
    if serviceName == "" {
        serviceName = "prayerProcessorService" // Default service name
    }

    exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerEndpoint)))
    if err != nil {
        log.Fatalf("Failed to initialize Jaeger exporter: %v", err)
    }

    tp := trace.NewTracerProvider(
        trace.WithBatcher(exp),
        trace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
        )),
    )

    otel.SetTracerProvider(tp)
}

func main() {
    initTracer()

    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    kafkaProducer, err := service.NewKafkaProducer(cfg.KafkaBrokers)
    if err != nil {
        log.Fatalf("Failed to initialize Kafka producer: %v", err)
    }

    app := fiber.New()

    // Register the Swagger handler route
    app.Get("/swagger/*", swagger.Handler) // swagger.Handler serves the Swagger UI

    app.Post("/processPrayer", api.HandlePrayerRequest(kafkaProducer))

    log.Fatal(app.Listen(":" + cfg.ServerPort))
}

