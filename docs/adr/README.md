# Architecture Decision Records

Формат — облегчённый [Michael Nygard](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions):
контекст → решение → последствия. Новый ADR — копия [template.md](template.md) со следующим номером.
Принятые ADR не переписываются: если решение меняется, создаётся новый ADR со статусом `Supersedes NNNN`.

ADR 0001–0010 зафиксированы ретроспективно 2026-11-02 по уже существующему коду.

| № | Решение | Статус |
|---|---|---|
| [0001](0001-feature-based-layered-architecture.md) | Feature-based слоистая архитектура | Accepted |
| [0002](0002-interfaces-in-contract-files.md) | Интерфейсы в `contract.go` на стороне реализации | Accepted |
| [0003](0003-repository-implementation-in-postgres-subpackage.md) | Реализация репозитория в подпакете `postgres/` | Accepted |
| [0004](0004-raw-sql-with-pgx.md) | Сырой SQL через pgx, без ORM | Accepted |
| [0005](0005-optimistic-locking-with-version.md) | Оптимистичная блокировка через `version` | Accepted |
| [0006](0006-patch-with-nullable.md) | PATCH через `Nullable[T]` и `ApplyPatch` | Accepted |
| [0007](0007-sentinel-errors-to-http-status.md) | Sentinel-ошибки → HTTP-статус | Accepted |
| [0008](0008-stdlib-http-routing.md) | `net/http` ServeMux, версия API в префиксе | Accepted |
| [0009](0009-statistics-computed-on-read.md) | Статистика вычисляется при чтении | Accepted |
| [0010](0010-pagination-limits.md) | Обязательная пагинация с лимитами | Accepted |
