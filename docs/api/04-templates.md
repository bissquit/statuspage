# Шаблоны событий

API для управления шаблонами событий (инцидентов и плановых работ).

## Список шаблонов

**GET** `/api/v1/templates`

🔒 **Требует авторизации: operator**

Получение списка всех шаблонов.

### Query Parameters

- `type` (опционально) - фильтр по типу: `incident` или `maintenance`

### Response (200 OK)

```json
[
  {
    "id": "aa0e8400-e29b-41d4-a716-446655440000",
    "name": "Database Outage",
    "type": "incident",
    "title_template": "{{.ServiceName}} Database Unavailable",
    "message_template": "We are investigating reports of {{.ServiceName}} database being unavailable. Users may experience connection errors.",
    "impact": "major",
    "created_at": "2026-01-19T12:00:00Z",
    "updated_at": "2026-01-19T12:00:00Z"
  }
]
```

### Example

```bash
curl http://localhost:8080/api/v1/templates \
  -H "Authorization: Bearer $TOKEN"

curl http://localhost:8080/api/v1/templates?type=incident \
  -H "Authorization: Bearer $TOKEN"
```

---

## Получение шаблона

**GET** `/api/v1/templates/{id}`

🔒 **Требует авторизации: operator**

Получение шаблона по ID.

### Response (200 OK)

```json
{
  "id": "aa0e8400-e29b-41d4-a716-446655440000",
  "name": "Database Outage",
  "type": "incident",
  "title_template": "{{.ServiceName}} Database Unavailable",
  "message_template": "We are investigating reports of {{.ServiceName}} database being unavailable. Users may experience connection errors.",
  "impact": "major",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:00:00Z"
}
```

### Errors

- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - шаблон не найден

### Example

```bash
curl http://localhost:8080/api/v1/templates/aa0e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer $TOKEN"
```

---

## Создание шаблона

**POST** `/api/v1/templates`

🔒 **Требует авторизации: admin**

Создание нового шаблона.

### Request

```json
{
  "name": "Database Outage",
  "type": "incident",
  "title_template": "{{.ServiceName}} Database Unavailable",
  "message_template": "We are investigating reports of {{.ServiceName}} database being unavailable. Users may experience connection errors.",
  "impact": "major"
}
```

**Поля:**
- `name` (обязательное) - название шаблона
- `type` (обязательное) - тип: `incident` или `maintenance`
- `title_template` (обязательное) - Go template для заголовка
- `message_template` (обязательное) - Go template для сообщения
- `impact` (обязательное) - уровень воздействия

### Доступные переменные в шаблонах

- `{{.ServiceName}}` - название сервиса
- `{{.ServiceSlug}}` - slug сервиса
- `{{.Timestamp}}` - текущая дата/время
- `{{.CustomVar}}` - любая пользовательская переменная

### Response (201 Created)

```json
{
  "id": "aa0e8400-e29b-41d4-a716-446655440000",
  "name": "Database Outage",
  "type": "incident",
  "title_template": "{{.ServiceName}} Database Unavailable",
  "message_template": "We are investigating reports of {{.ServiceName}} database being unavailable. Users may experience connection errors.",
  "impact": "major",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:00:00Z"
}
```

### Errors

- `400` - некорректный JSON, валидация не пройдена или ошибка в синтаксисе template
- `401` - требуется авторизация
- `403` - недостаточно прав (требуется роль admin)

### Example

```bash
curl -X POST http://localhost:8080/api/v1/templates \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Database Outage",
    "type": "incident",
    "title_template": "{{.ServiceName}} Database Unavailable",
    "message_template": "We are investigating reports of {{.ServiceName}} database being unavailable.",
    "impact": "major"
  }'
```

---

## Обновление шаблона

**PATCH** `/api/v1/templates/{id}`

🔒 **Требует авторизации: admin**

Обновление существующего шаблона.

### Request

```json
{
  "name": "Database Outage (Updated)",
  "message_template": "We are experiencing issues with {{.ServiceName}}. Our team is working on a resolution."
}
```

**Все поля опциональные:**
- `name` - новое название
- `title_template` - новый шаблон заголовка
- `message_template` - новый шаблон сообщения
- `impact` - новый уровень воздействия

### Response (200 OK)

```json
{
  "id": "aa0e8400-e29b-41d4-a716-446655440000",
  "name": "Database Outage (Updated)",
  "type": "incident",
  "title_template": "{{.ServiceName}} Database Unavailable",
  "message_template": "We are experiencing issues with {{.ServiceName}}. Our team is working on a resolution.",
  "impact": "major",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:05:00Z"
}
```

### Errors

- `400` - некорректный JSON или ошибка в синтаксисе template
- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - шаблон не найден

### Example

```bash
curl -X PATCH http://localhost:8080/api/v1/templates/aa0e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "impact": "critical"
  }'
```

---

## Удаление шаблона

**DELETE** `/api/v1/templates/{id}`

🔒 **Требует авторизации: admin**

Удаление шаблона.

### Response (204 No Content)

### Errors

- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - шаблон не найден

### Example

```bash
curl -X DELETE http://localhost:8080/api/v1/templates/aa0e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer $TOKEN"
```

---

## Рендеринг шаблона (preview)

**POST** `/api/v1/templates/{id}/render`

🔒 **Требует авторизации: operator**

Предварительный просмотр шаблона с переменными.

### Request

```json
{
  "variables": {
    "ServiceName": "User Database",
    "ServiceSlug": "user-db"
  }
}
```

### Response (200 OK)

```json
{
  "title": "User Database Database Unavailable",
  "message": "We are investigating reports of User Database database being unavailable. Users may experience connection errors."
}
```

### Errors

- `400` - некорректный JSON или ошибка рендеринга
- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - шаблон не найден

### Example

```bash
curl -X POST http://localhost:8080/api/v1/templates/aa0e8400-e29b-41d4-a716-446655440000/render \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "variables": {
      "ServiceName": "Payment Service"
    }
  }'
```

---

## Создание события из шаблона

**POST** `/api/v1/events/from-template`

🔒 **Требует авторизации: operator**

Создание события на основе шаблона.

### Request

```json
{
  "template_id": "aa0e8400-e29b-41d4-a716-446655440000",
  "variables": {
    "ServiceName": "User Database"
  },
  "status": "investigating",
  "service_ids": ["550e8400-e29b-41d4-a716-446655440000"]
}
```

**Поля:**
- `template_id` (обязательное) - ID шаблона
- `variables` (обязательное) - переменные для подстановки
- `status` (обязательное) - начальный статус события
- `service_ids` (обязательное) - затронутые сервисы
- `started_at` (опционально) - время начала

### Response (201 Created)

```json
{
  "id": "770e8400-e29b-41d4-a716-446655440000",
  "type": "incident",
  "title": "User Database Database Unavailable",
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

- `400` - некорректный JSON, валидация не пройдена или ошибка рендеринга шаблона
- `401` - требуется авторизация
- `403` - недостаточно прав
- `404` - шаблон не найден

### Example

```bash
curl -X POST http://localhost:8080/api/v1/events/from-template \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "template_id": "aa0e8400-e29b-41d4-a716-446655440000",
    "variables": {
      "ServiceName": "Payment API"
    },
    "status": "investigating",
    "service_ids": ["550e8400-e29b-41d4-a716-446655440000"]
  }'
```

---

## Примеры шаблонов

### Инцидент: Degraded Performance

```json
{
  "name": "Degraded Performance",
  "type": "incident",
  "title_template": "{{.ServiceName}} Experiencing Performance Issues",
  "message_template": "We are seeing increased latency on {{.ServiceName}}. Response times are currently {{.Latency}}. We are investigating the cause.",
  "impact": "minor"
}
```

### Плановые работы: Database Migration

```json
{
  "name": "Database Migration",
  "type": "maintenance",
  "title_template": "Scheduled Maintenance: {{.ServiceName}}",
  "message_template": "We will be performing a database migration for {{.ServiceName}} on {{.ScheduledDate}}. Expected downtime: {{.Duration}}.",
  "impact": "major"
}
```

### Инцидент: Security Incident

```json
{
  "name": "Security Incident",
  "type": "incident",
  "title_template": "Security Alert: {{.ServiceName}}",
  "message_template": "We have detected unusual activity on {{.ServiceName}}. As a precaution, we have temporarily disabled the service while we investigate. Your data remains secure.",
  "impact": "critical"
}
```
