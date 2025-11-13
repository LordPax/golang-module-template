# Stage 1: Build the application
FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    go mod download && \
    swag init && \
    go build

# Stage 2: Create the final lightweight image
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/golang-api .
COPY --from=builder /app/clouds.yaml .
EXPOSE 8080

CMD ["./golang-api"]
