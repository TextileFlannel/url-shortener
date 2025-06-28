FROM golang:1.24.0 AS builder

RUN apt-get update && apt-get install -y protobuf-compiler

RUN CGO_ENABLED=0 go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && CGO_ENABLED=0 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && CGO_ENABLED=0 go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /app
COPY . .

RUN go mod download \
    && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/link.proto \
    && CGO_ENABLED=0 GOOS=linux go build -o /url-shortener ./cmd/server

FROM alpine:latest

RUN apk add --no-cache bash postgresql-client

COPY --from=builder /url-shortener /url-shortener
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY ./migrations /migrations
COPY .env .

ENV DATABASE_URL="postgres://postgres:postgres@postgres:5432/links?sslmode=disable"

EXPOSE 8080 50051

CMD ["sh", "-c", "goose -dir /migrations postgres \"$DATABASE_URL\" up && /url-shortener"]