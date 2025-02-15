# Default Service

### Русский язык 🇷🇺

Это репозиторий с реализованным boilerplate-кодом. Разработчику достаточно просто склонировать репозиторий и поменять название модуля, чтобы начать писать бизнес-логику своего приложения.

Требуется:

1) Docker
2) docker-compose
3) Make

Основные команды:

1) `make help` - получить справку о командах
2) `make compose-up-prod` - запустить сервис (для продакшена)
3) `make compose-up-dev` - запустить только Postgres, RabbitMQ и т.п.
4) `make run-app` - используется вместе с `make compose-up-dev`. Запускает само приложение (для разработки)
5) `make install-deps` - установить инструменты (для разработки)
6) `make migrate-create name="migration_name"` - создать миграцию

Структура:

```
├── Dockerfile
├── Makefile
├── .env               # для частоизменяемых или приватных переменных
├── README.md
├── cmd
│   └── app
│       └── main.go    # точка запуска приложения 
├── config
│   ├── config.go      # реализация конфига приложения
│   └── config.yml     # файл для остальных переменных
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal           # содержит только бизнес-логику приложения
│   ├── app
│   │   ├── app.go     # собираем все слои и запускаем бизнес-логику
│   │   └── migrate.go
│   ├── controller     # всевозможные контроллеры для HTTP, GRPC, RabbitMQ и т.Д
│   ├── entity         # сущности приложения
│   └── usecase        # обработка контроллеров
├── migrations         # миграции проекта
└── pkg                # папка с вспомогательным кодом
    ├── httpserver
    ├── logger
    └── postgres

```

