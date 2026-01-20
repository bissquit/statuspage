# Публичный статус

🌐 Все эндпоинты в этом разделе **публичные** — не требуют авторизации.

## Список сервисов

**GET** `/api/v1/services`

Получение списка всех сервисов с их текущим статусом.

### Response (200 OK)

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

### Example

```bash
curl http://localhost:8080/api/v1/services
```

---

## Получение сервиса

**GET** `/api/v1/services/{slug}`

Получение информации о конкретном сервисе по slug.

### Response (200 OK)

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

### Errors

- `404` - сервис не найден

### Example

```bash
curl http://localhost:8080/api/v1/services/api-gateway
```

---

## Список групп

**GET** `/api/v1/groups`

Получение списка всех групп сервисов.

### Response (200 OK)

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

### Example

```bash
curl http://localhost:8080/api/v1/groups
```

---

## Получение группы

**GET** `/api/v1/groups/{slug}`

Получение информации о конкретной группе по slug.

### Response (200 OK)

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

### Errors

- `404` - группа не найдена

### Example

```bash
curl http://localhost:8080/api/v1/groups/core-services
```

---

## Список событий

**GET** `/api/v1/events`

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

### Example

```bash
# Все события
curl http://localhost:8080/api/v1/events

# Только инциденты
curl http://localhost:8080/api/v1/events?type=incident

# Только активные инциденты
curl "http://localhost:8080/api/v1/events?type=incident&status=investigating"
```

---

## Получение события

**GET** `/api/v1/events/{id}`

Получение полной информации о событии с историей обновлений.

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
      "message": "The issue has been resolved.",
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

## Health Check

**GET** `/healthz`

Проверка работоспособности приложения (liveness probe).

### Response (200 OK)

```
OK
```

### Example

```bash
curl http://localhost:8080/healthz
```

---

## Readiness Check

**GET** `/readyz`

Проверка готовности приложения принимать запросы (readiness probe).
Проверяет подключение к базе данных.

### Response (200 OK)

```
OK
```

### Response (503 Service Unavailable)

```
Database unavailable
```

### Example

```bash
curl http://localhost:8080/readyz
```

---

## Примеры использования для Status Page

### Простая Status Page (HTML + JavaScript)

```html
<!DOCTYPE html>
<html>
<head>
    <title>System Status</title>
    <style>
        .operational { color: green; }
        .degraded_performance { color: orange; }
        .partial_outage { color: red; }
        .major_outage { color: darkred; }
        .under_maintenance { color: blue; }
    </style>
</head>
<body>
    <h1>System Status</h1>
    <div id="services"></div>
    <h2>Recent Incidents</h2>
    <div id="incidents"></div>

    <script>
        const API_BASE = 'http://localhost:8080/api/v1';

        async function loadServices() {
            const response = await fetch(`${API_BASE}/services`);
            const services = await response.json();
            
            const html = services.map(s => `
                <div class="${s.status}">
                    <strong>${s.name}</strong>: ${s.status}
                </div>
            `).join('');
            
            document.getElementById('services').innerHTML = html;
        }

        async function loadIncidents() {
            const response = await fetch(`${API_BASE}/events?type=incident`);
            const incidents = await response.json();
            
            const html = incidents.map(i => `
                <div>
                    <h3>${i.title} (${i.status})</h3>
                    <p>Impact: ${i.impact}</p>
                    <p>Started: ${new Date(i.started_at).toLocaleString()}</p>
                </div>
            `).join('');
            
            document.getElementById('incidents').innerHTML = html;
        }

        loadServices();
        loadIncidents();
        
        // Refresh every 60 seconds
        setInterval(() => {
            loadServices();
            loadIncidents();
        }, 60000);
    </script>
</body>
</html>
```

### Dashboard с деталями события

```bash
#!/bin/bash

# Получить все сервисы и их статусы
echo "=== System Status ==="
curl -s http://localhost:8080/api/v1/services | jq -r '.[] | "\(.name): \(.status)"'

echo -e "\n=== Active Incidents ==="
INCIDENTS=$(curl -s "http://localhost:8080/api/v1/events?type=incident")
echo "$INCIDENTS" | jq -r '.[] | select(.status != "resolved") | "\(.title) [\(.impact)] - \(.status)"'

echo -e "\n=== Scheduled Maintenance ==="
MAINTENANCE=$(curl -s "http://localhost:8080/api/v1/events?type=maintenance")
echo "$MAINTENANCE" | jq -r '.[] | select(.status != "completed") | "\(.title) - scheduled for \(.scheduled_for)"'
```

### Monitoring скрипт

```bash
#!/bin/bash

# Проверка всех сервисов и отправка алерта если есть проблемы
SERVICES=$(curl -s http://localhost:8080/api/v1/services)

ISSUES=$(echo "$SERVICES" | jq -r '.[] | select(.status != "operational") | "\(.name): \(.status)"')

if [ ! -z "$ISSUES" ]; then
    echo "⚠️  Service issues detected:"
    echo "$ISSUES"
    # Здесь можно добавить отправку алерта в Slack, email и т.д.
else
    echo "✅ All services operational"
fi
```

---

## Интеграция с мониторингом

### Prometheus metrics endpoint

В будущих версиях планируется добавить `/metrics` эндпоинт для Prometheus с метриками:

- `statuspage_service_status{service="name", status="operational"}` - текущий статус сервисов
- `statuspage_incidents_total{impact="major"}` - количество инцидентов
- `statuspage_incident_duration_seconds{service="name"}` - длительность инцидентов

### Webhook уведомления

Для интеграции с внешними системами можно использовать polling публичных эндпоинтов или настроить подписки через API уведомлений.
