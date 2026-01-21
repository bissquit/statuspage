# События

API для управления инцидентами и плановыми работами.

## Получение токенов для работы

```bash
# Operator токен (для создания/обновления/просмотра событий)
OPERATOR_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "operator@example.com", "password": "admin123"}' | jq -r '.data.tokens.access_token')

echo "Operator token: $OPERATOR_TOKEN"

# Admin токен (для удаления событий)
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@example.com", "password": "admin123"}' | jq -r '.data.tokens.access_token')

echo "Admin token: $ADMIN_TOKEN"
```

## Список событий

**GET** `/api/v1/events`

🔒 **Требует авторизации: operator**

Получение списка всех событий (инцидентов и плановых работ).

### Query Parameters

- `type` (опционально) - фильтр по типу: `incident` или `maintenance`
- `status` (опционально) - фильтр по статусу

### Response (200 OK)

```json
[
  {
    "id": "770e8400-e29b-41d4-a716-446655440000",
    "type": "incident",
    "title": "API Gateway Downtime",
    "status": "investigating",
    "severity": "major",
    "service_ids": ["550e8400-e29b-41d4-a716-446655440000"],
    "started_at": "2026-01-19T12:00:00Z",
    "resolved_at": null,
    "created_at": "2026-01-19T12:00:00Z",
    "updated_at": "2026-01-19T12:00:00Z"
  }
]
```

**Типы событий:**
- `incident` - инцидент (незапланированный сбой)
- `maintenance` - плановые работы

**Статусы инцидентов:**
- `investigating` - расследуется
- `identified` - причина установлена
- `monitoring` - под наблюдением
- `resolved` - решено

**Статусы плановых работ:**
- `scheduled` - запланировано
- `in_progress` - в процессе
- `completed` - завершено

**Уровни серьёзности (severity):**
- `minor` - минимальное воздействие
- `major` - значительное воздействие
- `critical` - критическое воздействие

### Example

```bash
# Получить operator токен (см. начало документа)
OPERATOR_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "operator@example.com", "password": "admin123"}' | jq -r '.data.tokens.access_token')

curl http://localhost:8080/api/v1/events \
  -H "Authorization: Bearer $OPERATOR_TOKEN" | jq

curl "http://localhost:8080/api/v1/events?type=incident" \
  -H "Authorization: Bearer $OPERATOR_TOKEN" | jq

curl "http://localhost:8080/api/v1/events?status=investigating" \
  -H "Authorization: Bearer $OPERATOR_TOKEN" | jq
```

---

## Получение события

**GET** `/api/v1/events/{id}`

🔒 **Требует авторизации: operator**

Получение события по ID с полной историей обновлений.

### Response (200 OK)

```json
{
  "id": "770e8400-e29b-41d4-a716-446655440000",
  "type": "incident",
  "title": "API Gateway Downtime",
  "status": "resolved",
  "severity": "major",
  "service_ids": ["550e8400-e29b-41d4-a716-446655440000"],
  "started_at": "2026-01-19T12:00:00Z",
  "resolved_at": "2026-01-19T13:00:00Z",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T13:00:00Z",
  "updates": [
    {
      "id": "880e8400-e29b-41d4-a716-446655440000",
      "event_id": "770e8400-e29b-41d4-a716-446655440000",
      "status": "investigating",
      "message": "We are investigating reports of API Gateway being unavailable.",
      "created_at": "2026-01-19T12:00:00Z"
    },
    {
      "id": "990e8400-e29b-41d4-a716-446655440000",
      "event_id": "770e8400-e29b-41d4-a716-446655440000",
      "status": "resolved",
      "message": "The issue has been resolved. All services are operational.",
      "created_at": "2026-01-19T13:00:00Z"
    }
  ]
}
```

### Errors

- `404` - событие не найдено

### Example

```bash
curl http://localhost:8080/api/v1/events/770e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer $OPERATOR_TOKEN" | jq
```

---

## Создание события

**POST** `/api/v1/events`

🔒 **Требует авторизации: operator**

Создание нового события (инцидент или плановые работы).

### Request

```json
{
  "type": "incident",
  "title": "API Gateway Downtime",
  "description": "We are investigating reports of API Gateway being unavailable.",
  "status": "investigating",
  "severity": "major",
  "service_ids": ["550e8400-e29b-41d4-a716-446655440000"],
  "started_at": "2026-01-19T12:00:00Z",
  "notify_subscribers": true
}
```

**Поля:**
- `type` (обязательное) - тип события: `incident` или `maintenance`
- `title` (обязательное) - заголовок события
- `description` (обязательное) - описание события
- `status` (обязательное) - начальный статус
- `severity` (опционально) - уровень серьёзности: `minor`, `major`, `critical`
- `service_ids` (опционально) - массив ID затронутых сервисов
- `started_at` (опционально) - время начала (по умолчанию текущее время)
- `scheduled_start_at` (для maintenance) - запланированное время начала
- `scheduled_end_at` (для maintenance) - запланированное время окончания
- `notify_subscribers` (опционально) - отправить уведомления подписчикам

### Response (201 Created)

```json
{
  "id": "770e8400-e29b-41d4-a716-446655440000",
  "title": "API Gateway Downtime",
  "type": "incident",
  "status": "investigating",
  "severity": "major",
  "description": "We are investigating reports of API Gateway being unavailable.",
  "started_at": "2026-01-19T12:00:00Z",
  "resolved_at": null,
  "scheduled_start_at": null,
  "scheduled_end_at": null,
  "notify_subscribers": false,
  "template_id": null,
  "created_by": "550e8400-e29b-41d4-a716-446655440001",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:00:00Z",
  "service_ids": ["550e8400-e29b-41d4-a716-446655440000"]
}
```

### Errors

- `400` - некорректный JSON или валидация не пройдена
- `401` - требуется авторизация
- `403` - недостаточно прав (требуется роль operator)

### Example

```bash
# Получить OPERATOR_TOKEN (если ещё не получен)
OPERATOR_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "operator@example.com", "password": "admin123"}' | jq -r '.data.tokens.access_token')

# Создать инцидент
curl -X POST http://localhost:8080/api/v1/events \
  -H "Authorization: Bearer $OPERATOR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "incident",
    "title": "API Gateway Downtime",
    "description": "We are investigating the issue.",
    "status": "investigating",
    "severity": "major",
    "service_ids": ["9dc4217c-3354-4075-bc8b-b69b8febcea1"]
  }' | jq
```

---

## Добавление обновления к событию

**POST** `/api/v1/events/{id}/updates`

🔒 **Требует авторизации: operator**

Добавление нового обновления к существующему событию.

### Request

```json
{
  "status": "identified",
  "message": "The root cause has been identified. We are working on a fix."
}
```

**Поля:**
- `status` (обязательное) - новый статус события
- `message` (обязательное) - сообщение об обновлении

### Response (201 Created)

```json
{
  "data": {
    "id": "880e8400-e29b-41d4-a716-446655440000",
    "event_id": "770e8400-e29b-41d4-a716-446655440000",
    "status": "identified",
    "message": "The root cause has been identified. We are working on a fix.",
    "created_at": "2026-01-19T12:30:00Z"
  }
}
```

### Errors

- `400` - некорректный JSON или валидация не пройдена
- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - событие не найдено

### Example

```bash
curl -X POST http://localhost:8080/api/v1/events/770e8400-e29b-41d4-a716-446655440000/updates \
  -H "Authorization: Bearer $OPERATOR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "resolved",
    "message": "The issue has been resolved."
  }' | jq
```

---

## Обновление события

**PATCH** `/api/v1/events/{id}`

🔒 **Требует авторизации: operator**

Обновление метаданных события (без добавления обновления в timeline).

### Request

```json
{
  "title": "API Gateway Partial Outage",
  "severity": "minor"
}
```

**Все поля опциональные:**
- `title` - новый заголовок
- `severity` - новый уровень серьёзности
- `service_ids` - новый список затронутых сервисов

### Response (200 OK)

```json
{
  "id": "770e8400-e29b-41d4-a716-446655440000",
  "type": "incident",
  "title": "API Gateway Partial Outage",
  "status": "investigating",
  "severity": "minor",
  "service_ids": ["550e8400-e29b-41d4-a716-446655440000"],
  "started_at": "2026-01-19T12:00:00Z",
  "resolved_at": null,
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:35:00Z"
}
```

### Errors

- `400` - некорректный JSON или валидация не пройдена
- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - событие не найдено

### Example

```bash
curl -X PATCH http://localhost:8080/api/v1/events/770e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer $OPERATOR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "severity": "minor"
  }' | jq
```

---

## Удаление события

**DELETE** `/api/v1/events/{id}`

🔒 **Требует авторизации: admin**

Удаление события.

### Response (204 No Content)

### Errors

- `401` - требуется авторизация
- `403` - недостаточно прав (требуется роль admin)
- `404` - событие не найдено

### Example

```bash
# Удаление требует admin роли
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@example.com", "password": "admin123"}' | jq -r '.data.tokens.access_token')

curl -X DELETE http://localhost:8080/api/v1/events/770e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer $ADMIN_TOKEN" | jq
```

---

## Полный пример workflow инцидента

```bash
# Шаг 1: Получить токены
echo "=== Получение токенов ==="
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@example.com", "password": "admin123"}' | jq -r '.data.tokens.access_token')

OPERATOR_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "operator@example.com", "password": "admin123"}' | jq -r '.data.tokens.access_token')

echo "Токены получены успешно"

# Шаг 2: Создать сервис (требуется admin)
echo -e "\n=== Создание сервиса ==="
SERVICE=$(curl -s -X POST http://localhost:8080/api/v1/services \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Database Service",
    "slug": "database-service",
    "description": "Main database service"
  }')

echo "$SERVICE" | jq
SERVICE_ID=$(echo "$SERVICE" | jq -r '.data.id')

# Шаг 3: Создать инцидент
echo -e "\n=== Создание инцидента ==="
EVENT=$(curl -s -X POST http://localhost:8080/api/v1/events \
  -H "Authorization: Bearer $OPERATOR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "incident",
    "title": "Database Connection Issues",
    "description": "Users are experiencing connection timeouts to the database.",
    "status": "investigating",
    "severity": "major",
    "service_ids": ["'"$SERVICE_ID"'"]
  }')

echo "$EVENT" | jq
EVENT_ID=$(echo "$EVENT" | jq -r '.id')

# Шаг 4: Обновление 1 - Identified
echo -e "\n=== Обновление 1: Identified ==="
UPDATE1=$(curl -s -X POST http://localhost:8080/api/v1/events/$EVENT_ID/updates \
  -H "Authorization: Bearer $OPERATOR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "identified",
    "message": "The issue is caused by a database server running out of memory."
  }')

echo "$UPDATE1" | jq

# Шаг 5: Обновление 2 - Monitoring
echo -e "\n=== Обновление 2: Monitoring ==="
UPDATE2=$(curl -s -X POST http://localhost:8080/api/v1/events/$EVENT_ID/updates \
  -H "Authorization: Bearer $OPERATOR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "monitoring",
    "message": "Memory has been increased. Monitoring the situation."
  }')

echo "$UPDATE2" | jq

# Шаг 6: Обновление 3 - Resolved
echo -e "\n=== Обновление 3: Resolved ==="
UPDATE3=$(curl -s -X POST http://localhost:8080/api/v1/events/$EVENT_ID/updates \
  -H "Authorization: Bearer $OPERATOR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "resolved",
    "message": "Database is stable. All connections are working normally."
  }')

echo "$UPDATE3" | jq

# Шаг 7: Получение полной истории инцидента
echo -e "\n=== Полная история инцидента ==="
curl -s http://localhost:8080/api/v1/events/$EVENT_ID/updates \
  -H "Authorization: Bearer $OPERATOR_TOKEN" | jq

echo -e "\n✅ Workflow завершён успешно!"
```
