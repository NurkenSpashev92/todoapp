# Known issues / технический долг

Задокументировано, чтобы не «чинить» случайно. Исправлять — отдельной задачей.

## Опечатки в публичных идентификаторах

`UserSerivce`, `core_errors.ErrInvlidArgument`, `RegiserRoutes`, `RegiserAPIRoutes`,
`NewHTTPResponseHander`, `NewUsersHTTPHander`, файл `response/hander.go`, поле `Config.Databse`.
Переименование затрагивает весь проект — делать только целиком и по явному запросу.
В **новом** коде новые имена пиши без опечаток, но существующие идентификаторы используй как есть.

## Несоответствия

- Transport у `users` лежит в `features/users/transport/http/`, у `tasks` и `statistics` — в `features/<x>/transport/`. Для новых фич используй `transport/`.
- `tasks.description` в миграции `NOT NULL`, а в domain/DTO — опциональное. Создание задачи без описания упадёт с 500.
- Конфликт версии при PATCH: `users` возвращает `ErrNotFound` (404), `tasks` — `ErrConflict` (409).
- `JSONResponse` выставляет `Content-Type: text/plain; charset=utf-8` вместо `application/json`.
- `CreateTaskRequest` содержит `completed`, `created_at`, `completed_At`, которые игнорируются (`NewTaskUninitialized`).
- Ошибка FK-нарушения при создании задачи не оборачивает sentinel → 500 вместо 400/404.
- `docker-compose.yml` / `Dockerfile` запускают `air -c .air.toml`, но `.air.toml` нет в репозитории.

## Отсутствует

- Тесты (unit для domain/service на моках интерфейсов, integration для репозиториев).
- Линтер (`golangci-lint`) и CI.
- README.
