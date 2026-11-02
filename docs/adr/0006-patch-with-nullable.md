# 0006. PATCH через `Nullable[T]` и `ApplyPatch`

- Статус: Accepted
- Дата: 2026-11-02

## Контекст

В PATCH нужно различать «поле не передано» и «поле явно `null`».

## Решение

- `domain.Nullable[T]{Value *T; Set bool}`; HTTP-обёртка `core_http_types.Nullable[T]` с `UnmarshalJSON` выставляет `Set=true`, если ключ есть в JSON.
- `Patch<X>Request` (transport) → `domain.<X>Patch` через `ToDomain()`.
- `(*X).ApplyPatch(patch)`: валидирует патч, применяет к копии, валидирует копию, присваивает.
- Производные поля вычисляются в `ApplyPatch` (например, `completed_at` при смене `completed`).

## Последствия

- Бизнес-правила обновления живут в domain, а не в SQL.
- Сохраняется сущность целиком (все изменяемые колонки в `UPDATE`).
