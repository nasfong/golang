# Use a Go version that satisfies air's requirement (e.g., go 1.23)
FROM golang:1.23 AS builder

WORKDIR /app

# Copy the local go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Install Air for hot-reloading
RUN go install github.com/cosmtrek/air@v1.61.7

# Run the Go application
CMD ["air"]
