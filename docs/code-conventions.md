# Code conventions

Документ описывает, **как уже написан код** в проекте. Новый код должен выглядеть так же.
Эталоны: `features/users` и `features/tasks`.

## 1. Структура файлов

- Одна операция — один файл, имя `<verb>_<entity>.go`:
  `create_user.go`, `get_user.go` (по id), `get_users.go` (список), `patch_user.go`, `delete_user.go`.
  Одинаковые имена файлов во всех трёх слоях фичи.
- `contract.go` — интерфейсы слоя (в `repository/` и `service/`).
- `repository.go` / `service.go` / `transport.go` — struct + конструктор (+ `Routes()` в transport).
- `model.go` / `models.go` — Model-структуры БД и `<x>DomainsFromModels`.
- `dto_common.go` — общий response DTO фичи и `<x>DTOFromDomain`, `<x>sDTOFromDomain`.

## 2. Пакеты и импорты

Имена пакетов — `snake_case` с указанием фичи и слоя (не совпадают с именем каталога):

| Каталог | Пакет |
|---|---|
| `core/errors` | `core_errors` |
| `core/logger` | `core_logger` |
| `core/repository/postgres/pool` | `core_postgres_pool` |
| `core/transport/http/request` | `core_http_request` |
| `core/transport/http/response` | `core_http_response` |
| `core/transport/http/server` | `core_http_server` |
| `core/transport/http/middleware` | `core_http_middleware` |
| `core/transport/http/types` | `core_http_types` |
| `features/<x>/repository` | `<x>_repository` |
| `features/<x>/repository/postgres` | `<x>_postgres_repository` |
| `features/<x>/service` | `<x>_service` |
| `features/<x>/transport` (у users — `transport/http`) | `<x>_transport_http` |
| `core/domain` | `domain` |

Импорт всегда с явным алиасом, равным имени пакета (кроме `domain`):

```go
import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_errors "github.com/nurkenspashev92/todoapp/internal/core/errors"
)
```

Группы: stdlib → внешние → внутренние, разделены пустой строкой.

## 3. Именование

- Структуры слоёв: `UserRepository`, `TaskService`, `TasksHTTPHandler`.
- Интерфейсы: `<X>RepositoryInterface`, `<X>ServiceInterface`.
- Конструкторы: `New<Type>(deps) *<Type>`.
- Методы CRUD: `Create<X>`, `Get<X>s`, `Get<X>ById`, `Patch<X>ById`, `Delete<X>ById`.
- Суффикс `Id` (а не `ID`) в именах методов и параметров: `GetUserById`, `userId`, `AuthorUserId`. Поле `ID` в структурах — заглавными.
- Request/Response DTO: `Create<X>Request`, `Patch<X>Request`, `<X>DTOResponse`; ответы операций — именованные типы над DTO: `type GetTaskResponse TaskDTOResponse`.
- Ключи query-параметров — константы: `limitQueryParamsKey = "limit"`.
- JSON-поля — `snake_case`.

## 4. Domain (`internal/core/domain`)

- Модель + полный конструктор `New<X>(...)` + `New<X>Uninitialized(...)` для создания (ID/Version = `UninitializedID`/`UninitializedVersion`).
- `func (x *X) Validate() error` — инварианты; ошибки оборачивают `core_errors.ErrInvlidArgument`.
- Nullable-поля — указатели (`*string`, `*time.Time`).
- Частичное обновление: `<X>Patch` с полями `Nullable[T]`, `New<X>Patch`, `(p *<X>Patch) Validate()`, `(x *X) ApplyPatch(patch)`.
  `ApplyPatch` работает с копией (`tmp := *x`), валидирует её и только потом присваивает `*x = tmp`.
- Предкомпилированные значения (regexp и т.п.) — переменные уровня пакета.

## 5. Repository (`features/<x>/repository/postgres`)

Шаблон метода:

```go
func (r *TaskRepository) GetTaskById(ctx context.Context, id int) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, ...
	FROM todoapp.tasks
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var taskModel TaskModel

	err := row.Scan(&taskModel.ID, ...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("task not found by id=%d: %w", id, core_errors.ErrNotFound)
		}

		return domain.Task{}, fmt.Errorf("failed to scan task: %w", err)
	}

	return domain.NewTask(taskModel.ID, ...), nil
}
```

Правила:
- Зависимость — `core_postgres_pool.Pool` (интерфейс), не `*pgxpool.Pool`.
- SQL в raw-строке с отступом в одну табуляцию, ключевые слова заглавными, схема `todoapp.` всегда явно, точка с запятой в конце.
- `INSERT/UPDATE ... RETURNING <все колонки>` — возвращаем актуальную строку из БД.
- Сканируем в `<X>Model`, потом конвертируем в domain через `domain.New<X>(...)`.
- Списки: `defer rows.Close()` сразу после проверки ошибки `Query`, срез с `make(..., 0, capacity)`, проверка `rows.Err()` после цикла.
- Опциональные фильтры — в SQL: `WHERE ($1::INTEGER IS NULL OR author_user_id=$1)`.
- `DELETE` — через `Exec`, `RowsAffected() == 0` → `ErrNotFound`.
- `UPDATE` — `WHERE id=$n AND version=$m`, `version=version+1`; нет строки → конфликт версии.

## 6. Service (`features/<x>/service`)

- Валидирует аргументы, которые не валидирует domain: `limit`/`offset`, `user_id > 0`, `from < to`.
- Ошибки валидации — `fmt.Errorf("...: %w", core_errors.ErrInvlidArgument)`.
- Перед сохранением вызывает `domain.Validate()` / `ApplyPatch()`.
- Сценарий PATCH: `GetById` → `ApplyPatch` → `PatchById` (версия из прочитанной сущности).
- Пагинация: `limit == nil` → `domain.DefaultLimit`; `limit > domain.MaxLimit` → 400.
- Каждая ошибка репозитория оборачивается с контекстом: `fmt.Errorf("faild to get task from repository: %w", err)`.

## 7. Transport (`features/<x>/transport`)

Шаблон хендлера:

```go
func (h *TasksHTTPHandler) GetTaskById(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHander(log, rw)

	taskId, err := getTaskIdPathValue(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task id from path",
		)
		return
	}

	taskDomain, err := h.taskService.GetTaskById(ctx, taskId)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task in service",
		)
		return
	}

	response := GetTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
```

Правила:
- Логгер — только `core_logger.FromContext(ctx)`.
- Тело запроса — `core_http_request.DecodeAndValidate(r, &request)`:
  если у request есть метод `Validate() error`, вызывается он, иначе — теги `validate:"..."` (go-playground/validator).
- Path/query — `core_http_request.GetIntPathValue`, `GetIntQueryParam`, `GetDateQueryParam` (формат `YYYY-MM-DD`).
- Статусы: `POST` → 201, `GET`/`PATCH` → 200, `DELETE` → 204 (`NoContentResponse`).
- DTO никогда не уходит в service — только domain.
- Маршруты описываются в `Routes() []core_http_server.Route` с паттерном Go 1.22+ (`/tasks/{id}`).

## 8. Ошибки

- Всегда оборачивай: `fmt.Errorf("<что делали>: %w", err)`. Сообщения — с маленькой буквы, без точки.
- Классификация только через sentinel-ошибки `core_errors.ErrNotFound | ErrInvlidArgument | ErrConflict`.
- Не логируй ошибку в service/repository — её логирует `ErrorResponse` один раз.
- Значения в сообщениях: `id=%d`, имена полей в обратных кавычках: `` invalide `Title` len: %d ``.

## 9. Форматирование

- `gofmt` обязателен.
- Если вызов не помещается в строку или у него больше 2–3 аргументов — каждый аргумент на своей строке, закрывающая скобка на отдельной строке (см. `ErrorResponse(...)`, `domain.NewTask(...)`).
- Пустая строка перед `return` после блока, после `defer cancel()`, между логическими шагами.
- Комментарии не пишутся для очевидного кода.

## 10. Конфигурация и логирование

- Конфиг пакета: `Config` с тегами `envconfig`, `NewConfig()`, `NewConfigMust()`.
- Новые переменные окружения — сразу в `.env.example` с рабочим значением.
- Логи: `zap` с полями (`zap.String`, `zap.Error`), уровень по смыслу (Debug — шаги инициализации/трассировка).

## 11. Миграции

- `make migrate-create seq=<snake_name>` → пара `NNNNNN_<name>.up.sql` / `.down.sql`.
- Применённые миграции не редактируются.
- Ограничения дублируются в БД (`CHECK`) и в `domain.Validate()`.
- Для колонок фильтрации/внешних ключей создаётся индекс.
