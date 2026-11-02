# API

Базовый префикс: `/api/v1`. Формат — JSON. Ошибки: `{"message": "...", "error": "..."}`.

## Users

| Метод | Путь | Описание | Успех |
|---|---|---|---|
| POST | `/users` | Создать: `{"full_name", "phone_number"?}` | 201 |
| GET | `/users?limit=&offset=` | Список (limit по умолчанию 20, максимум 100) | 200 |
| GET | `/users/{id}` | Получить по id | 200 |
| PATCH | `/users/{id}` | Частичное обновление `full_name`, `phone_number` | 200 |
| DELETE | `/users/{id}` | Удалить | 204 |

## Tasks

| Метод | Путь | Описание | Успех |
|---|---|---|---|
| POST | `/tasks` | Создать: `{"title", "description"?, "author_user_id"}` | 201 |
| GET | `/tasks?user_id=&limit=&offset=` | Список, фильтр по `author_user_id` | 200 |
| GET | `/tasks/{id}` | Получить по id | 200 |
| PATCH | `/tasks/{id}` | `title`, `description`, `completed` (при `true` ставится `completed_at`, при `false` — сбрасывается) | 200 |
| DELETE | `/tasks/{id}` | Удалить | 204 |

PATCH: отсутствующее поле — не меняется, `null` — сбросить (для nullable-полей). Параллельное изменение → 409 (tasks).

## Statistics

| Метод | Путь | Описание |
|---|---|---|
| GET | `/statistics?user_id=&from=YYYY-MM-DD&to=YYYY-MM-DD` | Агрегаты по задачам, период `[from, to)` по `created_at` |

```json
{
  "tasks_created": 10,
  "tasks_completed": 4,
  "tasks_completed_rate": 40,
  "tasks_average_completion_time": "3h25m10s"
}
```
`tasks_completed_rate` и `tasks_average_completion_time` равны `null`, если данных нет.
