# Start from the official Golang base image
FROM golang:1.22.1-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies. 
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server/main.go

# Start a new stage from scratch
FROM alpine:latest  

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy pages, templates and static folders
COPY --from=builder /app/pages ./pages
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
