# Архитектура

## Слои

| Слой | Пакет | Ответственность | Знает о |
|---|---|---|---|
| Domain | `internal/core/domain` | Модели (`User`, `Task`, `Statistics`), конструкторы, `Validate()`, Patch-типы, `ApplyPatch()` | `core/errors` |
| Transport | `features/<x>/transport` | Разбор HTTP (path/query/body), DTO ↔ domain, выбор статуса ответа | `<X>ServiceInterface`, domain, `core/transport/http/*` |
| Service | `features/<x>/service` | Бизнес-правила: валидация аргументов, пагинация, сценарии (get → patch → save) | `<X>RepositoryInterface`, domain, `core/errors` |
| Repository contract | `features/<x>/repository` | Только `contract.go` с интерфейсом | domain |
| Repository impl | `features/<x>/repository/postgres` | SQL, Model ↔ domain, маппинг `pgx.ErrNoRows` → `ErrNotFound` | `core_postgres_pool.Pool`, domain |

Слой не импортирует слой выше себя. Фичи не импортируют друг друга.

## Поток запроса

```
HTTP request
  → middleware: RequestID → Logger (кладёт logger в ctx) → Trace → Panic
  → APIVersionRouter (/api/v1, ServeMux "METHOD /path")
  → <X>HTTPHandler.<Op>
      core_http_request.GetIntPathValue / GetIntQueryParam / DecodeAndValidate
      DTO → domain
  → <X>Service.<Op>        (валидация, бизнес-логика)
  → <X>Repository.<Op>     (context.WithTimeout, SQL, Scan в Model, Model → domain)
  ← domain
  ← core_http_response.JSONResponse / NoContentResponse / ErrorResponse
```

## Wiring (`cmd/todoapp/main.go`)

Ручная инъекция зависимостей через конструкторы, в порядке:

```go
logger.Debug("initializing feature", zap.String("feature", "tasks"))
taskRepository := tasks_postgres_repository.NewTaskRepository(pool)
taskService := tasks_service.NewTaskService(taskRepository)
tasksTransportHTTP := tasks_transport_http.NewTaskHTTPHandler(taskService)
...
apiVersionRouter.RegiserRoutes(tasksTransportHTTP.Routes()...)
```

## Ошибки → HTTP

Определено в `core/transport/http/response/hander.go`:

| Ошибка (`errors.Is`) | Статус | Уровень лога |
|---|---|---|
| `core_errors.ErrInvlidArgument` | 400 | Warn |
| `core_errors.ErrNotFound` | 404 | Debug |
| `core_errors.ErrConflict` | 409 | Warn |
| любая другая | 500 | Error |
| panic | 500 | Error |

Тело ошибки: `{"message": "<msg хендлера>", "error": "<цепочка err.Error()>"}`.

## Конфигурация

Через переменные окружения (`kelseyhightower/envconfig`), префиксы:
`POSTGRES_*` (pool), `HTTP_*` (server), `LOGGER_*` (logger). Пример — `.env.example`.
Каждый пакет конфигурации: `Config`, `NewConfig() (Config, error)`, `NewConfigMust() Config` (panic при ошибке).

## База данных

- Схема `todoapp`, таблицы `users`, `tasks` (`tasks.author_user_id → users.id`).
- Колонка `version` для оптимистичной блокировки.
- Миграции — `golang-migrate` в контейнере `todo-migrate`.
