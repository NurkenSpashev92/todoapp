# Добавление фичи или эндпоинта

## Новая операция в существующей фиче (пример: `tasks`)

1. **Domain** — при необходимости новые поля/методы в `internal/core/domain/task.go` (+ `Validate`).
2. **Repository**
   - метод в `features/tasks/repository/contract.go` → `TaskRepositoryInterface`;
   - файл `features/tasks/repository/postgres/<verb>_task.go` по шаблону из `code-conventions.md` §5.
3. **Service**
   - метод в `features/tasks/service/contract.go` → `TaskServiceInterface`;
   - файл `features/tasks/service/<verb>_task.go`.
4. **Transport**
   - файл `features/tasks/transport/<verb>_task.go` (request/response DTO + хендлер);
   - маршрут в `Routes()` в `transport.go`.
5. **Миграция**, если меняется схема: `make migrate-create seq=<name>`.
6. **Документация**: `docs/api.md`.
7. Проверка: `gofmt -l . && go build ./... && go vet ./...`.

Compile-time проверки `var _ ...Interface = (*...)(nil)` сразу покажут, если интерфейс и реализация разошлись.

## Новая фича (пример: `projects`)

```
internal/features/projects/
├── repository/
│   ├── contract.go              # package projects_repository
│   └── postgres/                # package projects_postgres_repository
│       ├── repository.go        # ProjectRepository, NewProjectRepository, var _ assertion
│       ├── model.go             # ProjectModel, projectDomainsFromModels
│       └── <verb>_project.go
├── service/                     # package projects_service
│   ├── contract.go              # ProjectServiceInterface + var _ assertion
│   ├── service.go               # ProjectService, NewProjectService
│   └── <verb>_project.go
└── transport/                   # package projects_transport_http
    ├── transport.go             # ProjectsHTTPHandler, NewProjectHTTPHandler, Routes()
    ├── dto_common.go            # ProjectDTOResponse, projectDTOFromDomain
    └── <verb>_project.go
```

Плюс:
- `internal/core/domain/project.go` — модель, конструкторы, `Validate`, `ProjectPatch`, `ApplyPatch`;
- миграция с таблицей в схеме `todoapp`, `version BIGINT NOT NULL DEFAULT 1`, `CHECK`-ограничениями и индексами;
- wiring в `cmd/todoapp/main.go`: `logger.Debug("initializing feature", ...)`, repository → service → handler, `apiVersionRouter.RegiserRoutes(...)`.

## Чек-лист ревью

- [ ] Интерфейсы в `contract.go`, реализация репозитория в `repository/postgres/`
- [ ] Один файл — одна операция, имена файлов совпадают между слоями
- [ ] `context.WithTimeout(ctx, r.pool.OpTimeout())` в каждом методе репозитория
- [ ] `defer rows.Close()` после `Query`
- [ ] `pgx.ErrNoRows` → `ErrNotFound`, конфликт версии → `ErrConflict`, плохой ввод → `ErrInvlidArgument`
- [ ] Ошибки обёрнуты через `%w`
- [ ] Списки ограничены `limit` (default/max)
- [ ] DTO не проникает в service
- [ ] Новые env-переменные в `.env.example`
- [ ] `gofmt`, `go build`, `go vet` чистые
