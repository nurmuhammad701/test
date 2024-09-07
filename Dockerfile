# Stage 1: Build stage
FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .


# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]