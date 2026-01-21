# Каталог сервисов

API для управления сервисами и группами сервисов.

## Получение токенов для работы

```bash
# Admin токен (для CRUD сервисов и групп)
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@example.com", "password": "admin123"}' | jq -r '.data.tokens.access_token')

echo "Admin token: $ADMIN_TOKEN"
```

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
curl http://localhost:8080/api/v1/services | jq
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
curl http://localhost:8080/api/v1/services/api-gateway | jq
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
# Получить admin токен (см. начало документа)
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@example.com", "password": "admin123"}' | jq -r '.data.tokens.access_token')

curl -X POST http://localhost:8080/api/v1/services \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "API Gateway",
    "slug": "api-gateway",
    "description": "Основной API шлюз"
  }' | jq
```

---

### Обновление сервиса

**PATCH** `/api/v1/services/{slug}`

🔒 **Требует авторизации: admin**

Обновление существующего сервиса.

#### Request

```json
{
  "name": "API Gateway",
  "slug": "api-gateway",
  "status": "degraded_performance"
}
```

**ВАЖНО: Все поля обязательные** (это особенность текущей реализации):
- `name` - название сервиса
- `slug` - slug сервиса (должен совпадать с URL параметром)
- `status` - статус сервиса
- `description` (опционально) - описание
- `group_id` (опционально) - ID группы

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
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "API Gateway",
    "slug": "api-gateway",
    "status": "operational"
  }' | jq
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
  -H "Authorization: Bearer $ADMIN_TOKEN" | jq
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
curl http://localhost:8080/api/v1/groups | jq
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
curl http://localhost:8080/api/v1/groups/core-services | jq
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
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Core Services",
    "slug": "core-services",
    "description": "Основные сервисы платформы"
  }' | jq
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
# Обновление группы также требует все поля
curl -X PATCH http://localhost:8080/api/v1/groups/core-services \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Core Services",
    "slug": "core-services",
    "description": "Новое описание"
  }' | jq
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
  -H "Authorization: Bearer $ADMIN_TOKEN" | jq
```

---

## Полный пример workflow

```bash
# Шаг 1: Получить admin токен
echo "=== Получение admin токена ==="
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }' | jq -r '.data.tokens.access_token')

echo "Токен получен: ${ADMIN_TOKEN:0:20}..."

# Шаг 2: Создать группу сервисов
echo -e "\n=== Создание группы сервисов ==="
GROUP=$(curl -s -X POST http://localhost:8080/api/v1/groups \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Backend Services",
    "slug": "backend-services",
    "description": "Все backend сервисы"
  }')

echo "$GROUP" | jq
GROUP_ID=$(echo "$GROUP" | jq -r '.data.id')

# Шаг 3: Создать первый сервис
echo -e "\n=== Создание первого сервиса ==="
SERVICE1=$(curl -s -X POST http://localhost:8080/api/v1/services \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Authentication API",
    "slug": "auth-api",
    "description": "API для аутентификации пользователей",
    "group_id": "'"$GROUP_ID"'"
  }')

echo "$SERVICE1" | jq
SERVICE1_ID=$(echo "$SERVICE1" | jq -r '.data.id')

# Шаг 4: Создать второй сервис
echo -e "\n=== Создание второго сервиса ==="
SERVICE2=$(curl -s -X POST http://localhost:8080/api/v1/services \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Payment Gateway",
    "slug": "payment-gateway",
    "description": "Сервис обработки платежей",
    "group_id": "'"$GROUP_ID"'"
  }')

echo "$SERVICE2" | jq
SERVICE2_ID=$(echo "$SERVICE2" | jq -r '.data.id')

# Шаг 5: Просмотр всех сервисов (публичный эндпоинт)
echo -e "\n=== Список всех сервисов ==="
curl -s http://localhost:8080/api/v1/services | jq

# Шаг 6: Обновить статус первого сервиса
echo -e "\n=== Обновление статуса сервиса ==="
curl -s -X PATCH http://localhost:8080/api/v1/services/auth-api \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Authentication API",
    "slug": "auth-api",
    "status": "degraded"
  }' | jq

# Шаг 7: Просмотр конкретного сервиса
echo -e "\n=== Просмотр сервиса ==="
curl -s http://localhost:8080/api/v1/services/auth-api | jq

# Шаг 8: Просмотр группы со всеми сервисами
echo -e "\n=== Просмотр группы ==="
curl -s http://localhost:8080/api/v1/groups/backend-services | jq

echo -e "\n✅ Workflow завершён успешно!"
```
