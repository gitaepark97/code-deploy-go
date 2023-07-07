# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8000
CMD ["/app/main"]