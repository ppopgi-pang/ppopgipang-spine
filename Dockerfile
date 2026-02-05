FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o spine-app ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/spine-app .
EXPOSE 8080
CMD ["./spine-app"]
