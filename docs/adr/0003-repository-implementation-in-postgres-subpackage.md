# 0003. Реализация репозитория в подпакете `postgres/`

- Статус: Accepted
- Дата: 2026-11-02

## Контекст

Контракт репозитория не должен зависеть от конкретного хранилища.

## Решение

- `features/<x>/repository/` содержит только `contract.go` (пакет `<x>_repository`).
- Реализация на PostgreSQL — `features/<x>/repository/postgres/` (пакет `<x>_postgres_repository`).
- Service зависит только от `<x>_repository`; `main.go` импортирует `.../repository/postgres`.

## Последствия

- Другую реализацию (in-memory для тестов, кэш) можно положить рядом в `repository/<impl>/`.
- Пакет `postgres` импортирует контракт ради compile-time проверки — обратной зависимости нет.
