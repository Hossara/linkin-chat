FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app cmd/main.go

FROM alpine:latest

WORKDIR /app/bin
COPY --from=builder /go/bin/app /app/bin

RUN chmod +x /app/bin/app

CMD ["/app/bin/app", "--config", "/app/bin/config.json"]