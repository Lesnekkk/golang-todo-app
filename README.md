# Golang Todo App

REST API приложение на Go — учебный проект. Реализует управление пользователями, задачами и статистику.

## Технологический стек

| Компонент         | Технология                        |
|-------------------|-----------------------------------|
| Язык              | Go 1.25+                          |
| HTTP-роутер       | `gorilla/mux`                     |
| База данных       | PostgreSQL                        |
| Драйвер БД        | `jackc/pgx/v5`                    |
| Кэш               | Redis                             |
| Клиент Redis      | `go-redis/v9`                     |
| Логгер            | `go.uber.org/zap`                 |
| Конфигурация      | `kelseyhightower/envconfig`       |
| Миграции БД       | `golang-migrate`                  |
| Деплой            | Docker / Docker Compose           |

---

## Архитектура

Проект следует принципам **чистой архитектуры** (Clean Architecture).
Каждая фича (`users`, `tasks`, `statistics`, `web`) разделена на три слоя:

```
Transport (HTTP Handler)  ←  декодирует запрос, вызывает сервис, формирует ответ
      ↓
Service (Business Logic)  ←  валидация, доменная логика
      ↓
Repository (Data Access)  ←  SQL-запросы к PostgreSQL / Redis-кэш

Domain (Core)             ←  сущности и инварианты, без зависимостей
```

**Cache-aside (Redis)**: при GET сначала ищем в Redis, при промахе берём из PostgreSQL и кладём в кэш (TTL 5 минут). При PATCH и DELETE запись в кэше инвалидируется.

**Dependency Inversion (DIP)**: интерфейсы определяются в потребляющем слое:
- `UserRepository` живёт в `users_service`
- `UsersService` живёт в `users_transport_http`

---

## Структура проекта

```
.
├── cmd/todoapp/          # Точка входа, Dockerfile
├── internal/
│   ├── core/             # Общие компоненты
│   │   ├── config/       # Конфигурация приложения
│   │   ├── domain/       # Сущности: User, Task, Statistics, Nullable, File
│   │   ├── errors/       # Sentinel-ошибки: ErrNotFound, ErrConflict, ErrInvalidArgument
│   │   ├── logger/       # Zap-логгер
│   │   ├── repository/postgres/  # Пул соединений pgx
│   │   ├── repository/redis/     # Redis-клиент
│   │   └── transport/http/       # Middleware, Nullable HTTP-тип
│   └── features/         # Бизнес-фичи
│       ├── users/        # CRUD пользователей (Postgres + Redis)
│       ├── tasks/        # CRUD задач (Postgres + Redis)
│       ├── statistics/   # Статистика
│       └── web/          # Раздача статических HTML-страниц
├── migrations/           # SQL-миграции
├── public/               # Статические файлы (index.html)
├── docker-compose.yaml
├── Makefile
└── postman_collection.json
```

---

## Локальный запуск

```bash
# 1. Скопировать .env файл
cp .env.example .env

# 2. Заполнить переменные окружения
nano .env

# 3. Поднять PostgreSQL
make env-up

# 4. Открыть порты для локального доступа
make env-port-forward

# 5. Применить миграции
make migrate-up

# 6. Запустить приложение в Docker
make app-up
```

После запуска:
- Главная страница: `http://localhost:8080/`
- API: `http://localhost:8080/api/v1/`

---

## API

### Пользователи `/api/v1/users`

| Метод    | Путь           | Описание                         |
|----------|----------------|----------------------------------|
| `POST`   | `/users`       | Создать пользователя             |
| `GET`    | `/users`       | Список пользователей (пагинация) |
| `GET`    | `/users/{id}`  | Получить пользователя по ID      |
| `PATCH`  | `/users/{id}`  | Частично обновить пользователя   |
| `DELETE` | `/users/{id}`  | Удалить пользователя             |

### Задачи `/api/v1/tasks`

| Метод    | Путь           | Описание                                          |
|----------|----------------|---------------------------------------------------|
| `POST`   | `/tasks`       | Создать задачу (поле `author_user_id` обязательно)|
| `GET`    | `/tasks`       | Список задач (фильтр `?user_id=`, пагинация)      |
| `GET`    | `/tasks/{id}`  | Получить задачу по ID                             |
| `PATCH`  | `/tasks/{id}`  | Частично обновить задачу (Three-state logic)      |
| `DELETE` | `/tasks/{id}`  | Удалить задачу                                    |

### Статистика `/api/v1/statistics`

| Метод  | Путь            | Описание                                                     |
|--------|-----------------|--------------------------------------------------------------|
| `GET`  | `/statistics`   | Статистика задач (фильтры: `user_id`, `from`, `to` YYYY-MM-DD) |

---

## Переменные окружения

| Переменная        | Описание                   | Пример              |
|-------------------|----------------------------|---------------------|
| `HTTP_ADDR`       | Адрес HTTP-сервера         | `:8080`             |
| `POSTGRES_HOST`   | Хост PostgreSQL            | `localhost`         |
| `POSTGRES_PORT`   | Порт PostgreSQL            | `5432`              |
| `POSTGRES_USER`   | Пользователь БД            | `todoapp`           |
| `POSTGRES_PASSWORD` | Пароль БД              | `password`          |
| `POSTGRES_DB`     | Имя базы данных            | `todoapp`           |
| `REDIS_HOST`      | Хост Redis                 | `localhost`         |
| `REDIS_PORT`      | Порт Redis                 | `6379`              |
| `LOGGER_LEVEL`    | Уровень логирования        | `INFO`              |
| `LOGGER_FOLDER`   | Папка для лог-файлов       | `out/logs`          |
| `PROJECT_ROOT`    | Корень проекта (путь к public/) | `/app`         |
