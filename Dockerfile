# Use Golang as the base image
FROM golang:1.24 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (to leverage Docker cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Fetch dependencies
RUN go mod tidy

# Build the project
RUN go build -o /bot main.go

# Create a minimal runtime image
FROM golang:1.24 AS runner

# Set working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /bot /app/bot

# Run the Telegram bot
CMD ["/app/bot"]
