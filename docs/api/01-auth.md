# Аутентификация

API для регистрации и аутентификации пользователей.

## Регистрация

**POST** `/api/v1/auth/register`

Регистрация нового пользователя. По умолчанию создаётся с ролью `user`.

### Request

```json
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

### Response (201 Created)

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "user@example.com",
  "role": "user",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:00:00Z"
}
```

### Errors

- `400` - некорректный JSON или валидация не пройдена
- `409` - пользователь с таким email уже существует

### Example

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newuser@example.com",
    "password": "securepassword123"
  }' | jq
```

---

## Логин

**POST** `/api/v1/auth/login`

Аутентификация пользователя.

### Request

```json
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

### Response (200 OK)

```json
{
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "user@example.com",
      "role": "user",
      "created_at": "2026-01-19T12:00:00Z",
      "updated_at": "2026-01-19T12:00:00Z"
    },
    "tokens": {
      "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "expires_in": 900
    }
  }
}
```

**Важно:** сохраните `access_token` для последующих запросов и `refresh_token` для обновления токена.

### Errors

- `400` - некорректный JSON
- `401` - неверные учётные данные

### Example

```bash
# Логин с тестовым пользователем
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "user123"
  }' | jq -r '.data.tokens.access_token' > /tmp/token.txt

export TOKEN=$(cat /tmp/token.txt)

# Или с админом для выполнения административных операций
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }' | jq -r '.data.tokens.access_token' > /tmp/admin_token.txt

export ADMIN_TOKEN=$(cat /tmp/admin_token.txt)
```

---

## Обновление токена

**POST** `/api/v1/auth/refresh`

Обновление access токена с помощью refresh токена.

### Request

```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Response (200 OK)

```json
{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 900
  }
}
```

### Errors

- `400` - некорректный JSON
- `401` - недействительный refresh токен

### Example

```bash
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d "{
    \"refresh_token\": \"$REFRESH_TOKEN\"
  }" | jq -r '.data.access_token' > /tmp/token.txt

export TOKEN=$(cat /tmp/token.txt)
```

---

## Логаут

**POST** `/api/v1/auth/logout`

🔒 **Требует авторизации**

Выход из системы (инвалидация refresh токена).

### Request

```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Response (204 No Content)

### Errors

- `400` - некорректный JSON
- `401` - требуется авторизация

### Example

```bash
curl -X POST http://localhost:8080/api/v1/auth/logout \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
    \"refresh_token\": \"$REFRESH_TOKEN\"
  }" | jq
```

---

## Текущий пользователь

**GET** `/api/v1/me`

🔒 **Требует авторизации**

Получение информации о текущем пользователе.

### Response (200 OK)

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "user@example.com",
  "role": "user",
  "created_at": "2026-01-19T12:00:00Z",
  "updated_at": "2026-01-19T12:00:00Z"
}
```

### Errors

- `401` - требуется авторизация

### Example

```bash
curl http://localhost:8080/api/v1/me \
  -H "Authorization: Bearer $TOKEN" | jq
```

---

## Полный пример workflow

```bash
# Шаг 1: Регистрация нового пользователя
echo "=== Регистрация ==="
REGISTER_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newuser@example.com",
    "password": "securepass123"
  }')

echo "$REGISTER_RESPONSE" | jq

# Шаг 2: Логин с новым пользователем
echo -e "\n=== Логин ==="
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newuser@example.com",
    "password": "securepass123"
  }')

echo "$LOGIN_RESPONSE" | jq

TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.data.tokens.access_token')
REFRESH_TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.data.tokens.refresh_token')

# Шаг 3: Получение информации о текущем пользователе
echo -e "\n=== Получение текущего пользователя ==="
ME_RESPONSE=$(curl -s http://localhost:8080/api/v1/me \
  -H "Authorization: Bearer $TOKEN")

echo "$ME_RESPONSE" | jq

# Шаг 4: Обновление access токена
echo -e "\n=== Refresh токена ==="
REFRESH_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d "{\"refresh_token\":\"$REFRESH_TOKEN\"}")

echo "$REFRESH_RESPONSE" | jq

NEW_TOKEN=$(echo "$REFRESH_RESPONSE" | jq -r '.data.access_token')

# Шаг 5: Проверка нового токена
echo -e "\n=== Проверка нового токена ==="
curl -s http://localhost:8080/api/v1/me \
  -H "Authorization: Bearer $NEW_TOKEN" | jq

# Шаг 6: Логаут
echo -e "\n=== Логаут ==="
curl -s -X POST http://localhost:8080/api/v1/auth/logout \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $NEW_TOKEN" \
  -d "{\"refresh_token\":\"$REFRESH_TOKEN\"}" | jq

echo -e "\n✅ Workflow завершён успешно!"
```
