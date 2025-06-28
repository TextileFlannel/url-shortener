Структура проекта:

```
.
├── api
│   ├── link.pb.go
│   ├── link.proto
│   └── link_grpc.pb.go
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── handler
│   │   ├── create.go
│   │   └── get.go
│   ├── storage
│   │   ├── inmem.go
│   │   ├── inmem_test.go
│   │   ├── postgres.go
│   │   ├── postgres_test.go
│   │   ├── shortener.go
│   │   └── storage.go
│   └── service
│       ├── service.go
│       └── service_test.go
├── migrations
│   ├── 20250628211610_create_urls_table.sql
│   └── migrations.go
├── .env
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

# Клонируйте проект

```
git clone https://github.com/TextileFlannel/url-shortener
```

# Запуск сервера:

## Через docker:
```
make docker-compose-up
```

## Остановка:
```
make docker-compose-down
```

# Локально:

## Сборка проекта:
```
make build
```

## inmem:
```
make run
```

## postgres:
```
make run-postgres
```

## Генерация gRPC кода:
```
make generate
```

## Применение миграций:
```
migrate-up
```

## Запуск тестов:
```
make tests
```

# Примеры запросов:

## gRPC:

CreateLink:

```
grpcurl -plaintext -d '{"original_url": "https://danil.com"}' localhost:50051 api.LinkService/CreateLink
```

Пример ответа:

```
{
  "shortUrl": "JHt3oBNWn4"
}
```

GetLink :

```
grpcurl -plaintext -d '{"short_url": "JHt3oBNWn4"}' localhost:50051 api.LinkService/GetLink
```

Пример ответа:

```
{
  "originalUrl": "https://danil.com"
}
```

Пример ответа когда не надено:

```
ERROR:
  Code: Unknown
  Message: not found
```

## HTTP API:

POST:

```
curl -X POST -d "url=https://danil.com" http://localhost:8080/create
```

Пример ответа:

```fZKh7hZIeV```

GET:

```
curl http://localhost:8080/get/fZKh7hZIeV
```

Пример ответа:

```
<a href="https://danil.com">Found</a>
```
