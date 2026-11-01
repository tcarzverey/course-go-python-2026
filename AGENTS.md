# AGENTS.md

This file provides guidance to AI-agent (Codex, Claude, etc) when working with code in this repository.

## О репозитории

Учебный репозиторий курса «НИС Построение сервисов на Golang и Python». 

## Команды

Запуск из корня репозитория:

- `go test ./...`, `go build ./...` — все тесты/сборка по корневому модулю
- `go vet ./...`, `go fmt ./...` — статический анализ и форматирование
- `go mod tidy` — подтянуть зависимости

## Структура

- `materials/` — конспекты лекций
- `examples/lectureN/` — код с лекций по темам (defer, panic, errors, concurrency, cli, GraphQL, gRPC, websockets)
- `homeworks/hwN/` — домашние задания, по одному пакету на задание

## Схема домашнего задания

Каждое задание — пакет вида `homeworks/hwN/<task>/` с единообразным устройством:

- `contract.go` — фиксированные интерфейсы и сигнатуры (например, `AggregationResult`, `HttpClient` для url-агрегатора); менять их нельзя, правится только реализация в файлах-заготовках
- файлы-заготовки (`result.go`, `aggregator.go`, `handler.go`, `errors/your_error.go`) — здесь должно быть решение; к своей реализации можно добавлять методы, не входящие в интерфейс
- `README.md` — условие задачи (с примерами и ожидаемым форматом вывода)
- `questions.md` — вопросы по заданию, ответы пишутся прямо в файл
- публичные тесты в том же пакете (`*_test.go`, включая example-тесты вида `aggregator_example_test.go`); приватные тесты при проверке запускаются отдельно
