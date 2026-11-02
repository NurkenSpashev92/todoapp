# 0007. Sentinel-ошибки → HTTP-статус

- Статус: Accepted
- Дата: 2026-11-02

## Контекст

Слои ниже transport не должны знать про HTTP, но transport должен выбрать статус.

## Решение

- `core/errors`: `ErrNotFound`, `ErrInvlidArgument`, `ErrConflict`.
- Все слои оборачивают ошибки через `fmt.Errorf("...: %w", err)`.
- `HTTPResponseHandler.ErrorResponse` маппит через `errors.Is`: 400 / 404 / 409, иначе 500; логирует один раз с уровнем по типу.

## Последствия

- Новый класс ошибок = новая sentinel + ветка в `ErrorResponse`.
- Нельзя терять `%w` — иначе ошибка станет 500.
- Цепочка `err.Error()` уходит клиенту в поле `error` — не включай в сообщения секреты.
