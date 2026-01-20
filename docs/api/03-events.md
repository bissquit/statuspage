# События

API для управления инцидентами и плановыми работами.

## Список событий

**GET** `/api/v1/events` 🌐 **Публичный эндпоинт**

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
    "impact": "major",
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

**Уровни воздействия (impact):**
- `none` - нет воздействия
- `minor` - минимальное
- `major` - значительное
- `critical` - критическое

### Example

```bash
curl http://localhost:8080/api/v1/events
curl http://localhost:8080/api/v1/events?type=incident
curl http://localhost:8080/api/v1/events?status=investigating
```

---

## Получение события

**GET** `/api/v1/events/{id}` 🌐 **Публичный эндпоинт**

Получение события по ID с полной историей обновлений.

### Response (200 OK)

```json
{
  "id": "770e8400-e29b-41d4-a716-446655440000",
  "type": "incident",
  "title": "API Gateway Downtime",
  "status": "resolved",
  "impact": "major",
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
curl http://localhost:8080/api/v1/events/770e8400-e29b-41d4-a716-446655440000
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
  "message": "We are investigating reports of API Gateway being unavailable.",
  "status": "investigating",
  "impact": "major",
  "service_ids": ["550e8400-e29b-41d4-a716-446655440000"],
  "started_at": "2026-01-19T12:00:00Z",
  "scheduled_for": null
}
```

**Поля:**
- `type` (обязательное) - тип события: `incident` или `maintenance`
- `title` (обязательное) - заголовок события
- `message` (обязательное) - первоначальное сообщение
- `status` (обязательное) - начальный статус
- `impact` (обязательное) - уровень воздействия
- `service_ids` (обязательное) - массив ID затронутых сервисов
- `started_at` (опционально) - время начала (по умолчанию текущее время)
- `scheduled_for` (для maintenance) - запланированное время

### Response (201 Created)

```json
{
  "id": "770e8400-e29b-41d4-a716-446655440000",
  "type": "incident",
  "title": "API Gateway Downtime",
  "status": "investigating",
  "impact": "major",
  "service_ids": ["550e8400-e29b-41d4-a716-446655440000"],
  "started_at": "2026-01-19T12:00:00Z",
  "resolved_at": null,
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:00:00Z"
}
```

### Errors

- `400` - некорректный JSON или валидация не пройдена
- `401` - требуется авторизация
- `403` - недостаточно прав (требуется роль operator)

### Example

```bash
curl -X POST http://localhost:8080/api/v1/events \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "incident",
    "title": "API Gateway Downtime",
    "message": "We are investigating the issue.",
    "status": "investigating",
    "impact": "major",
    "service_ids": ["550e8400-e29b-41d4-a716-446655440000"]
  }'
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
  "id": "880e8400-e29b-41d4-a716-446655440000",
  "event_id": "770e8400-e29b-41d4-a716-446655440000",
  "status": "identified",
  "message": "The root cause has been identified. We are working on a fix.",
  "created_at": "2026-01-19T12:30:00Z"
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
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "resolved",
    "message": "The issue has been resolved."
  }'
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
  "impact": "minor"
}
```

**Все поля опциональные:**
- `title` - новый заголовок
- `impact` - новый уровень воздействия
- `service_ids` - новый список затронутых сервисов

### Response (200 OK)

```json
{
  "id": "770e8400-e29b-41d4-a716-446655440000",
  "type": "incident",
  "title": "API Gateway Partial Outage",
  "status": "investigating",
  "impact": "minor",
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
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "impact": "minor"
  }'
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
curl -X DELETE http://localhost:8080/api/v1/events/770e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer $TOKEN"
```

---

## Полный пример workflow инцидента

```bash
TOKEN="your_operator_token_here"

echo "=== Создание инцидента ==="
EVENT=$(curl -s -X POST http://localhost:8080/api/v1/events \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "incident",
    "title": "Database Connection Issues",
    "message": "Users are experiencing connection timeouts to the database.",
    "status": "investigating",
    "impact": "major",
    "service_ids": ["550e8400-e29b-41d4-a716-446655440000"]
  }')

EVENT_ID=$(echo $EVENT | jq -r '.id')
echo "Created event: $EVENT_ID"

echo -e "\n=== Обновление 1: Identified ==="
curl -X POST http://localhost:8080/api/v1/events/$EVENT_ID/updates \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "identified",
    "message": "The issue is caused by a database server running out of memory."
  }'

echo -e "\n\n=== Обновление 2: Monitoring ==="
curl -X POST http://localhost:8080/api/v1/events/$EVENT_ID/updates \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "monitoring",
    "message": "Memory has been increased. Monitoring the situation."
  }'

echo -e "\n\n=== Обновление 3: Resolved ==="
curl -X POST http://localhost:8080/api/v1/events/$EVENT_ID/updates \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "resolved",
    "message": "Database is stable. All connections are working normally."
  }'

echo -e "\n\n=== Получение полной истории ==="
curl http://localhost:8080/api/v1/events/$EVENT_ID | jq .
```
