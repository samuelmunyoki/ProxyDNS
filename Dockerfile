# Stage 1: Build the Go executable
FROM golang:1.16 AS builder

WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Install any needed dependencies
RUN go get -d -v ./...

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd/...

# Stage 2: Create a minimal image
FROM alpine:latest

WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /go/src/app/myapp .

# Set environment variables
ENV LOGFILE="logs/proxydns.log"
ENV LOGFILESIZE=50
ENV LOGFILEBACKUPS=3
ENV LOGFILEDAYS=7

# Run the executable
CMD ["./myapp"]
