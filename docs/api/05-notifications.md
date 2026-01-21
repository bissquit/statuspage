# Уведомления

API для управления каналами уведомлений и подписками.

## Получение токена для работы

```bash
# User токен (для управления своими каналами и подписками)
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "user123"}' | jq -r '.data.tokens.access_token')

echo "User token: $TOKEN"
```

## Каналы уведомлений

### Список каналов

**GET** `/api/v1/me/channels`

🔒 **Требует авторизации: user**

Получение списка каналов уведомлений текущего пользователя.

#### Response (200 OK)

```json
[
  {
    "id": "bb0e8400-e29b-41d4-a716-446655440000",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "type": "email",
    "target": "user@example.com",
    "is_enabled": true,
    "is_verified": true,
    "created_at": "2026-01-19T12:00:00Z",
    "updated_at": "2026-01-19T12:00:00Z"
  },
  {
    "id": "cc0e8400-e29b-41d4-a716-446655440000",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "type": "telegram",
    "target": "@username",
    "is_enabled": false,
    "is_verified": false,
    "created_at": "2026-01-19T12:05:00Z",
    "updated_at": "2026-01-19T12:05:00Z"
  }
]
```

**Типы каналов:**
- `email` - Email уведомления
- `telegram` - Telegram уведомления

#### Example

```bash
# Получить user токен (см. начало документа)
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "user123"}' | jq -r '.data.tokens.access_token')

curl http://localhost:8080/api/v1/me/channels \
  -H "Authorization: Bearer $TOKEN" | jq
```

---

### Создание канала

**POST** `/api/v1/me/channels`

🔒 **Требует авторизации: user**

Создание нового канала уведомлений.

#### Request

```json
{
  "type": "email",
  "target": "notifications@example.com"
}
```

**Поля:**
- `type` (обязательное) - тип канала: `email` или `telegram`
- `target` (обязательное) - адрес получателя (email или Telegram username)

**Примечание:** новый канал создаётся включённым (`is_enabled: true`), но не верифицированным (`is_verified: false`). Уведомления отправляются только на верифицированные каналы.

#### Response (201 Created)

```json
{
  "data": {
    "id": "bb0e8400-e29b-41d4-a716-446655440000",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "type": "email",
    "target": "notifications@example.com",
    "is_enabled": true,
    "is_verified": false,
    "created_at": "2026-01-19T12:00:00Z",
    "updated_at": "2026-01-19T12:00:00Z"
  }
}
```

#### Errors

- `400` - некорректный JSON или валидация не пройдена
- `401` - требуется авторизация

#### Example

```bash
curl -X POST http://localhost:8080/api/v1/me/channels \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "email",
    "target": "alerts@example.com"
  }' | jq
```

---

### Обновление канала

**PATCH** `/api/v1/me/channels/{id}`

🔒 **Требует авторизации: user**

Включение/отключение канала уведомлений.

#### Request

```json
{
  "is_enabled": false
}
```

#### Response (200 OK)

```json
{
  "id": "bb0e8400-e29b-41d4-a716-446655440000",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "type": "email",
  "target": "notifications@example.com",
  "is_enabled": false,
  "is_verified": true,
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:10:00Z"
}
```

#### Errors

- `400` - некорректный JSON
- `401` - требуется авторизация
- `403` - канал не принадлежит пользователю
- `404` - канал не найден

#### Example

```bash
curl -X PATCH http://localhost:8080/api/v1/me/channels/bb0e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "is_enabled": true
  }' | jq
```

---

### Верификация канала

**POST** `/api/v1/me/channels/{id}/verify`

🔒 **Требует авторизации: user**

Верификация канала уведомлений.

**Примечание:** в текущей реализации это упрощённая верификация. В production должна быть полноценная верификация с отправкой кода подтверждения.

#### Response (200 OK)

```json
{
  "id": "bb0e8400-e29b-41d4-a716-446655440000",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "type": "email",
  "target": "notifications@example.com",
  "is_enabled": true,
  "is_verified": true,
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:15:00Z"
}
```

#### Errors

- `401` - требуется авторизация
- `403` - канал не принадлежит пользователю
- `404` - канал не найден

#### Example

```bash
curl -X POST http://localhost:8080/api/v1/me/channels/bb0e8400-e29b-41d4-a716-446655440000/verify \
  -H "Authorization: Bearer $TOKEN" | jq
```

---

### Удаление канала

**DELETE** `/api/v1/me/channels/{id}`

🔒 **Требует авторизации: user**

Удаление канала уведомлений.

#### Response (204 No Content)

#### Errors

- `401` - требуется авторизация
- `403` - канал не принадлежит пользователю
- `404` - канал не найден

#### Example

```bash
curl -X DELETE http://localhost:8080/api/v1/me/channels/bb0e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer $TOKEN" | jq
```

---

## Подписки

### Получение подписки

**GET** `/api/v1/me/subscriptions`

🔒 **Требует авторизации: user**

Получение подписки текущего пользователя.

**Примечание:** подписка создаётся автоматически при первом обращении, если её не существует.

#### Response (200 OK)

```json
{
  "id": "dd0e8400-e29b-41d4-a716-446655440000",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "service_ids": [
    "550e8400-e29b-41d4-a716-446655440000",
    "660e8400-e29b-41d4-a716-446655440000"
  ],
  "created_at": "2026-01-19T12:00:00Z"
}
```

**Логика подписки:**
- Если `service_ids` пустой массив — подписка на все сервисы
- Если `service_ids` содержит ID — подписка только на указанные сервисы

#### Example

```bash
curl http://localhost:8080/api/v1/me/subscriptions \
  -H "Authorization: Bearer $TOKEN" | jq
```

---

### Создание/обновление подписки

**POST** `/api/v1/me/subscriptions`

🔒 **Требует авторизации: user**

Создание или обновление подписки пользователя.

#### Request для подписки на все сервисы

```json
{
  "service_ids": []
}
```

#### Request для подписки на конкретные сервисы

```json
{
  "service_ids": [
    "550e8400-e29b-41d4-a716-446655440000",
    "660e8400-e29b-41d4-a716-446655440000"
  ]
}
```

#### Response (200 OK)

```json
{
  "id": "dd0e8400-e29b-41d4-a716-446655440000",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "service_ids": [
    "550e8400-e29b-41d4-a716-446655440000"
  ],
  "created_at": "2026-01-19T12:00:00Z"
}
```

#### Errors

- `400` - некорректный JSON
- `401` - требуется авторизация

#### Example

```bash
# Подписаться на все сервисы
curl -X POST http://localhost:8080/api/v1/me/subscriptions \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "service_ids": []
  }' | jq

# Подписаться на конкретные сервисы
curl -X POST http://localhost:8080/api/v1/me/subscriptions \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "service_ids": ["550e8400-e29b-41d4-a716-446655440000"]
  }' | jq
```

---

### Удаление подписки

**DELETE** `/api/v1/me/subscriptions`

🔒 **Требует авторизации: user**

Удаление подписки пользователя (отписка от всех уведомлений).

#### Response (204 No Content)

#### Errors

- `401` - требуется авторизация
- `404` - подписка не найдена

#### Example

```bash
curl -X DELETE http://localhost:8080/api/v1/me/subscriptions \
  -H "Authorization: Bearer $TOKEN" | jq
```

---

## Полный пример workflow

```bash
# Получить user токен
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "user123"}' | jq -r '.data.tokens.access_token')

echo "=== 1. Создание Email канала ==="
EMAIL_CHANNEL=$(curl -s -X POST http://localhost:8080/api/v1/me/channels \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "email",
    "target": "alerts@example.com"
  }')

EMAIL_CHANNEL_ID=$(echo $EMAIL_CHANNEL | jq -r '.data.id')
echo "Created email channel: $EMAIL_CHANNEL_ID"

echo -e "\n=== 2. Верификация Email канала ==="
curl -X POST http://localhost:8080/api/v1/me/channels/$EMAIL_CHANNEL_ID/verify \
  -H "Authorization: Bearer $TOKEN" | jq

echo -e "\n\n=== 3. Создание Telegram канала ==="
TELEGRAM_CHANNEL=$(curl -s -X POST http://localhost:8080/api/v1/me/channels \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "telegram",
    "target": "@myusername"
  }')

TELEGRAM_CHANNEL_ID=$(echo $TELEGRAM_CHANNEL | jq -r '.data.id')

echo -e "\n=== 4. Верификация Telegram канала ==="
curl -X POST http://localhost:8080/api/v1/me/channels/$TELEGRAM_CHANNEL_ID/verify \
  -H "Authorization: Bearer $TOKEN" | jq

echo -e "\n\n=== 5. Список всех каналов ==="
curl http://localhost:8080/api/v1/me/channels \
  -H "Authorization: Bearer $TOKEN" | jq

echo -e "\n\n=== 6. Создание подписки на все сервисы ==="
curl -X POST http://localhost:8080/api/v1/me/subscriptions \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "service_ids": []
  }' | jq

echo -e "\n\n=== 7. Отключение Telegram канала ==="
curl -X PATCH http://localhost:8080/api/v1/me/channels/$TELEGRAM_CHANNEL_ID \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "is_enabled": false
  }' | jq

echo -e "\n\n=== 8. Получение текущей подписки ==="
curl http://localhost:8080/api/v1/me/subscriptions \
  -H "Authorization: Bearer $TOKEN" | jq
```

---

## Как работают уведомления

1. Пользователь создаёт один или несколько каналов (email, telegram)
2. Каждый канал должен быть верифицирован
3. Пользователь создаёт подписку на сервисы (все или конкретные)
4. При создании инцидента/обновления события система:
   - Находит всех пользователей, подписанных на затронутые сервисы
   - Для каждого пользователя проверяет включённые и верифицированные каналы
   - Отправляет уведомления через соответствующие каналы

**Важно:** уведомления отправляются только если канал:
- `is_enabled: true`
- `is_verified: true`
