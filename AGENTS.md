# AGENTS.md

Инструкции для AI-агентов и разработчиков, работающих с этим репозиторием.

## Проект

REST API трекера задач на Go + PostgreSQL.
Модуль: `github.com/nurkenspashev92/todoapp`, Go 1.26.

Фичи: `users` (CRUD), `tasks` (CRUD), `statistics` (read-only агрегация по задачам).

## Команды

```bash
make install            # go mod download + поднять postgres + миграции
make start              # postgres в docker, приложение через go run
make env-up / env-down  # поднять / остановить postgres
make migrate-create seq=<name>
make migrate-up / migrate-down

go build ./...          # сборка
go vet ./...            # статический анализ
gofmt -l .              # должно выводить пусто
```

Перед завершением любой задачи: `gofmt -l .` пусто, `go build ./...` и `go vet ./...` без ошибок.
Тестов в проекте пока нет — не утверждай, что изменения «протестированы», если не запускал их против БД.

## Архитектура (кратко)

```
cmd/todoapp/main.go                 # wiring: repository → service → transport → router
internal/core/                      # общий код, не зависит от фич
  domain/                           # доменные модели, валидация, Patch-типы
  errors/                           # sentinel-ошибки (ErrNotFound, ErrInvlidArgument, ErrConflict)
  logger/                           # zap-логгер, кладётся в context
  repository/postgres/pool/         # pgxpool + интерфейс Pool
  transport/http/                   # server, router, middleware, request/response хелперы
internal/features/<feature>/
  repository/contract.go            # <X>RepositoryInterface
  repository/postgres/              # реализация на pgx, один файл = одна операция
  service/contract.go               # <X>ServiceInterface
  service/                          # бизнес-логика, один файл = одна операция
  transport/                        # HTTP-хендлеры, DTO, Routes()
migrations/                         # golang-migrate, NNNNNN_name.{up,down}.sql
```

Поток зависимостей: `transport → service (interface) → repository (interface) → postgres`.
Все слои работают с типами из `internal/core/domain`.

## Ключевые правила

1. **Интерфейсы — в `contract.go`** в пакете, который их реализует:
   `<X>RepositoryInterface` в `features/<x>/repository/contract.go`,
   `<X>ServiceInterface` в `features/<x>/service/contract.go`.
   Рядом с реализацией — compile-time проверка `var _ Iface = (*Impl)(nil)`.
2. **Реализация репозитория — в `repository/postgres/`**, а не в корне `repository/`.
3. **Один файл — одна операция** во всех слоях: `create_task.go`, `get_task.go`, `get_tasks.go`, `patch_task.go`, `delete_task.go`.
4. **Ошибки** оборачивай через `fmt.Errorf("...: %w", err)`; для HTTP-статуса используй sentinel-ошибки из `core/errors`.
5. **SQL — сырой, через pgx**, схема `todoapp.`, плейсхолдеры `$1..$n`, `RETURNING` для получения строки.
6. **Каждый метод репозитория** начинается с `context.WithTimeout(ctx, r.pool.OpTimeout())`.
7. После `pool.Query` — всегда `defer rows.Close()`.
8. **PATCH** — через `core_http_types.Nullable[T]` → `domain.<X>Patch` → `ApplyPatch` + оптимистичная блокировка по `version`.
9. Списки — `limit`/`offset` (по умолчанию `domain.DefaultLimit`, максимум `domain.MaxLimit`).
10. Пиши код, как написан соседний: те же имена, алиасы импортов, форматирование многострочных вызовов. Комментариев в коде почти нет — не добавляй очевидных.

Подробно: [docs/code-conventions.md](docs/code-conventions.md).

## Не делай без запроса

- Не переименовывай существующие идентификаторы с опечатками (`UserSerivce`, `ErrInvlidArgument`, `RegiserRoutes`, `NewHTTPResponseHander`) — это ломает весь проект; см. [docs/known-issues.md](docs/known-issues.md).
- Не меняй уже применённые миграции — создавай новую.
- Не добавляй ORM, роутер-фреймворки и DI-контейнеры (см. ADR).
- Не коммить без явной просьбы.

## Документация

- [docs/architecture.md](docs/architecture.md) — слои, поток запроса, wiring
- [docs/code-conventions.md](docs/code-conventions.md) — код-стандарт
- [docs/adding-a-feature.md](docs/adding-a-feature.md) — пошаговый рецепт новой фичи / эндпоинта
- [docs/api.md](docs/api.md) — эндпоинты
- [docs/known-issues.md](docs/known-issues.md) — технический долг
- [docs/adr/](docs/adr/README.md) — архитектурные решения
