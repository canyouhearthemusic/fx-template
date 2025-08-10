# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app


# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN addgroup -g 1001 appgroup && \
    adduser -D -s /bin/sh -u 1001 -G appgroup appuser

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/.env* ./

RUN chown -R appuser:appgroup /root/

USER appuser

EXPOSE 8080

CMD ["./main"]