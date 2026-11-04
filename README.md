# Products REST API

Учебный проект на Go: REST API для работы с продуктами и подписками.

## Стек
- Go + Fiber
- PostgreSQL
- golang-migrate
- Swagger (swaggo)

## Запуск
```bash
createdb subscriptions
migrate -path migrations -database "postgres://postgres:1234@localhost:5432/subscriptions?sslmode=disable" up
go run .
