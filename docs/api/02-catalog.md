# Каталог сервисов

API для управления сервисами и группами сервисов.

## Сервисы

### Список сервисов

**GET** `/api/v1/services` 🌐 **Публичный эндпоинт**

Получение списка всех сервисов.

#### Response (200 OK)

```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "API Gateway",
    "slug": "api-gateway",
    "description": "Основной API шлюз",
    "status": "operational",
    "group_id": "660e8400-e29b-41d4-a716-446655440000",
    "created_at": "2026-01-19T12:00:00Z",
    "updated_at": "2026-01-19T12:00:00Z"
  }
]
```

**Статусы сервисов:**
- `operational` - работает нормально
- `degraded_performance` - снижена производительность
- `partial_outage` - частичный сбой
- `major_outage` - полный сбой
- `under_maintenance` - на обслуживании

#### Example

```bash
curl http://localhost:8080/api/v1/services
```

---

### Получение сервиса

**GET** `/api/v1/services/{slug}` 🌐 **Публичный эндпоинт**

Получение сервиса по slug.

#### Response (200 OK)

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "API Gateway",
  "slug": "api-gateway",
  "description": "Основной API шлюз",
  "status": "operational",
  "group_id": "660e8400-e29b-41d4-a716-446655440000",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:00:00Z"
}
```

#### Errors

- `404` - сервис не найден

#### Example

```bash
curl http://localhost:8080/api/v1/services/api-gateway
```

---

### Создание сервиса

**POST** `/api/v1/services`

🔒 **Требует авторизации: admin**

Создание нового сервиса.

#### Request

```json
{
  "name": "API Gateway",
  "slug": "api-gateway",
  "description": "Основной API шлюз",
  "group_id": "660e8400-e29b-41d4-a716-446655440000"
}
```

**Поля:**
- `name` (обязательное) - название сервиса
- `slug` (обязательное) - уникальный идентификатор (URL-friendly)
- `description` (опциональное) - описание
- `group_id` (опциональное) - ID группы

#### Response (201 Created)

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "API Gateway",
  "slug": "api-gateway",
  "description": "Основной API шлюз",
  "status": "operational",
  "group_id": "660e8400-e29b-41d4-a716-446655440000",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:00:00Z"
}
```

#### Errors

- `400` - некорректный JSON или валидация не пройдена
- `401` - требуется авторизация
- `403` - недостаточно прав (требуется роль admin)
- `409` - сервис с таким slug уже существует

#### Example

```bash
curl -X POST http://localhost:8080/api/v1/services \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "API Gateway",
    "slug": "api-gateway",
    "description": "Основной API шлюз"
  }'
```

---

### Обновление сервиса

**PATCH** `/api/v1/services/{slug}`

🔒 **Требует авторизации: admin**

Обновление существующего сервиса.

#### Request

```json
{
  "name": "API Gateway (Updated)",
  "description": "Обновлённое описание",
  "status": "degraded_performance"
}
```

**Все поля опциональные:**
- `name` - новое название
- `description` - новое описание
- `status` - новый статус
- `group_id` - новая группа

#### Response (200 OK)

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "API Gateway (Updated)",
  "slug": "api-gateway",
  "description": "Обновлённое описание",
  "status": "degraded_performance",
  "group_id": "660e8400-e29b-41d4-a716-446655440000",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:05:00Z"
}
```

#### Errors

- `400` - некорректный JSON или валидация не пройдена
- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - сервис не найден

#### Example

```bash
curl -X PATCH http://localhost:8080/api/v1/services/api-gateway \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "operational"
  }'
```

---

### Удаление сервиса

**DELETE** `/api/v1/services/{slug}`

🔒 **Требует авторизации: admin**

Удаление сервиса.

#### Response (204 No Content)

#### Errors

- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - сервис не найден

#### Example

```bash
curl -X DELETE http://localhost:8080/api/v1/services/api-gateway \
  -H "Authorization: Bearer $TOKEN"
```

---

## Группы сервисов

### Список групп

**GET** `/api/v1/groups` 🌐 **Публичный эндпоинт**

Получение списка всех групп.

#### Response (200 OK)

```json
[
  {
    "id": "660e8400-e29b-41d4-a716-446655440000",
    "name": "Core Services",
    "slug": "core-services",
    "description": "Основные сервисы платформы",
    "created_at": "2026-01-19T12:00:00Z",
    "updated_at": "2026-01-19T12:00:00Z"
  }
]
```

#### Example

```bash
curl http://localhost:8080/api/v1/groups
```

---

### Получение группы

**GET** `/api/v1/groups/{slug}` 🌐 **Публичный эндпоинт**

Получение группы по slug.

#### Response (200 OK)

```json
{
  "id": "660e8400-e29b-41d4-a716-446655440000",
  "name": "Core Services",
  "slug": "core-services",
  "description": "Основные сервисы платформы",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:00:00Z"
}
```

#### Errors

- `404` - группа не найдена

#### Example

```bash
curl http://localhost:8080/api/v1/groups/core-services
```

---

### Создание группы

**POST** `/api/v1/groups`

🔒 **Требует авторизации: admin**

Создание новой группы.

#### Request

```json
{
  "name": "Core Services",
  "slug": "core-services",
  "description": "Основные сервисы платформы"
}
```

#### Response (201 Created)

```json
{
  "id": "660e8400-e29b-41d4-a716-446655440000",
  "name": "Core Services",
  "slug": "core-services",
  "description": "Основные сервисы платформы",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:00:00Z"
}
```

#### Errors

- `400` - некорректный JSON или валидация не пройдена
- `401` - требуется авторизация
- `403` - недостаточно прав
- `409` - группа с таким slug уже существует

#### Example

```bash
curl -X POST http://localhost:8080/api/v1/groups \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Core Services",
    "slug": "core-services",
    "description": "Основные сервисы платформы"
  }'
```

---

### Обновление группы

**PATCH** `/api/v1/groups/{slug}`

🔒 **Требует авторизации: admin**

Обновление существующей группы.

#### Request

```json
{
  "name": "Core Services (Updated)",
  "description": "Обновлённое описание"
}
```

#### Response (200 OK)

```json
{
  "id": "660e8400-e29b-41d4-a716-446655440000",
  "name": "Core Services (Updated)",
  "slug": "core-services",
  "description": "Обновлённое описание",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:05:00Z"
}
```

#### Errors

- `400` - некорректный JSON или валидация не пройдена
- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - группа не найдена

#### Example

```bash
curl -X PATCH http://localhost:8080/api/v1/groups/core-services \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Новое описание"
  }'
```

---

### Удаление группы

**DELETE** `/api/v1/groups/{slug}`

🔒 **Требует авторизации: admin**

Удаление группы.

#### Response (204 No Content)

#### Errors

- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - группа не найдена

#### Example

```bash
curl -X DELETE http://localhost:8080/api/v1/groups/core-services \
  -H "Authorization: Bearer $TOKEN"
```
