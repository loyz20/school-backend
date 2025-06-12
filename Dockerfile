# Stage 1 - Build binary
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 👇 Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/school-backend

# Stage 2 - Run app in minimal image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .env

EXPOSE 8080

CMD ["./main"]
