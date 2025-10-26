# Stage - 1
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o loader main.go

# Stage - 2
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

WORKDIR /app

COPY --from=builder /app/loader .

COPY --from=builder /app/migrations ./migrations

RUN chown -R appuser:appuser /app

USER appuser

CMD ["./loader"]