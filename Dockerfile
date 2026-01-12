# Build stage
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./ # Uncomment when dependencies are added

RUN go mod download

COPY . .

RUN go build -o intellilog cmd/server/main.go

# Run stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/intellilog .
COPY --from=builder /app/web ./web

EXPOSE 8080

CMD ["./intellilog"]
