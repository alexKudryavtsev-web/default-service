# Default Service

### Languages:

- English
- Русский
<!-- - עִברִית -->

---

### English 🇬🇧

This is a repository with the implemented boilerplate code. A developer just needs to clone the repository and change the module name to start writing the business logic of their application.

Required:

1) Docker
2) docker-compose
3) Make

Main commands:

1) `make help` - get help on commands
2) `make compose-up-prod` - start app (for production)
3) `make compose-up-dev` - start only Postgres, RabbitMQ and other
4) `make run-app` - run the application itself (for development), used with `make compose-up-dev`
5) `make install-deps` - install tools (for development)
6) `make migrate-create name="migration_name"` - create migration
7) `make generate-docs` - generate Swagger-documentation. 

Documentation: /api/docs/index.html

Structure

```
├── Dockerfile
├── Makefile
├── .env               # for frequently changed or private variables
├── README.md
├── cmd
│   └── app
│       └── main.go    # app launch point 
├── config
│   ├── config.go      # parsing configs
│   └── config.yml     # for other variables
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal           # only the business logic of app
│   ├── app
│   │   ├── app.go     # connects all parts, launching app
│   │   └── migrate.go
│   ├── controller     # all kinds of controllers for HTTP, GRPC, RabbitMQ, etc.
│   ├── entity         # app entities
│   └── usecase        # controller processing
├── migrations         # migrations files
└── pkg                # only auxiliary code
    ├── httpserver
    ├── logger
    └── postgres
```

---

### Русский 🇷🇺

Это репозиторий с реализованным boilerplate-кодом. Разработчику достаточно просто склонировать репозиторий и поменять название модуля, чтобы начать писать бизнес-логику своего приложения.

Требуется:

1) Docker
2) docker-compose
3) Make

Основные команды:

1) `make help` - получить справку о командах
2) `make compose-up-prod` - запустить сервис (для продакшена)
3) `make compose-up-dev` - запустить только Postgres, RabbitMQ и т.п.
4) `make run-app` -  запусктить само приложение (для разработки), используется вместе с `make compose-up-dev`
5) `make install-deps` - установить инструменты (для разработки)
6) `make migrate-create name="migration_name"` - создать миграцию
6) `make generate-docs` - сгенерировать Swagger-документацию. 

Документация: /api/docs/index.html

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
│   ├── config.go      # парсинг конфигов
│   └── config.yml     # файл для остальных переменных
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal           # только бизнес-логику приложения
│   ├── app
│   │   ├── app.go     # сборка всех частей, запуск приложения
│   │   └── migrate.go
│   ├── controller     # всевозможные контроллеры для HTTP, GRPC, RabbitMQ и т.д.
│   ├── entity         # сущности приложения
│   └── usecase        # обработка контроллеров
├── migrations         # миграции
└── pkg                # только вспомогательный код
    ├── httpserver
    ├── logger
    └── postgres
```

<!-- ---

### עִברִית 🇮🇱

זהו מאגר עם קוד boilerplate ממומש. מפתח יכול פשוט לשכפל את המאגר ולשנות את שם המודול כדי להתחיל לכתוב את הלוגיקה העסקית של היישום שלו.

דרישות:

1) Docker
2) docker-compose
3) Make

פקודות עיקריות:

1) make help - לקבלת עזרה על הפקודות
2) make compose-up-prod - להפעלת השירות (לסביבת ייצור)
3) make compose-up-dev - להפעלת Postgres, RabbitMQ וכו' בלבד
4) make run-app - הפעלת היישום עצמו (לפיתוח), משמש בשילוב עם make compose-up-dev
5) make install-deps - להתקנת כלים (לפיתוח)
6) make migrate-create name="migration_name" - ליצירת מיגרציה

מבנה:

```
├── Dockerfile
├── Makefile
├── .env               # למשתנים שמשתנים לעתים קרובות או משתנים פרטיים
├── README.md
├── cmd
│  └── app
│    └── main.go       # נקודת התחלה של היישום
├── config
│  ├── config.go       # ניתוח תצורה (parsing)
│  └── config.yml      # קובץ למשתנים אחרים
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal           # רק הלוגיקה העסקית של היישום
│  ├── app
│  │  ├── app.go       # הרכבת כל החלקים, הפעלת היישום
│  │  └── migrate.go
│  ├── controller      # כל מיני בקרים (controllers) עבור HTTP, GRPC, RabbitMQ וכו'.
│  ├── entity          # ישויות (entities) של היישום
│  └── usecase         # טיפול בבקרים
├── migrations         # מיגרציות
└── pkg                # רק קוד עזר
  ├── httpserver
  ├── logger
  └── postgres
``` -->