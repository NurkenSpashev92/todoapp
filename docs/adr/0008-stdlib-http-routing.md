# 0008. `net/http` ServeMux, версия API в префиксе

- Статус: Accepted
- Дата: 2026-11-02

## Контекст

С Go 1.22 `http.ServeMux` поддерживает методы и path-параметры (`GET /tasks/{id}`).

## Решение

- Только стандартная библиотека: без gin/echo/chi.
- `APIVersionRouter` на каждую версию (`/api/v1`), фичи отдают `Routes() []core_http_server.Route`.
- Middleware — `func(http.Handler) http.Handler`, цепочка `ChainMiddleware`; глобальные: RequestID, Logger, Trace, Panic.
- Логгер с `request_id` передаётся через `context`.

## Последствия

- Минимум зависимостей; path-параметры — `r.PathValue`.
- Хелперы разбора запроса (`core_http_request`) и ответа (`core_http_response`) пишем сами.
