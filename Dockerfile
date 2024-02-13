# Build stage
FROM golang:1.18 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Vet the application - go vet to identify suspicious constructs
RUN go vet ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prayerProcessor .

# Runtime stage
FROM registry.access.redhat.com/ubi8/ubi-minimal

# Install ca-certificates
RUN microdnf update && microdnf install ca-certificates && microdnf clean all

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/prayerProcessor .

# Command to run the executable
CMD ["./prayerProcessor"]

