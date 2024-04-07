# Base image
FROM golang:1.22

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Set up the entry point
CMD ["./main"]