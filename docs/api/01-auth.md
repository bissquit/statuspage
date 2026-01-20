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
    "email": "user@example.com",
    "password": "securepassword123"
  }'
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
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "role": "user",
    "created_at": "2026-01-19T12:00:00Z",
    "updated_at": "2026-01-19T12:00:00Z"
  },
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2026-01-19T12:15:00Z"
}
```

**Важно:** сохраните `access_token` для последующих запросов и `refresh_token` для обновления токена.

### Errors

- `400` - некорректный JSON
- `401` - неверные учётные данные

### Example

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword123"
  }' | jq -r '.access_token' > /tmp/token.txt

export TOKEN=$(cat /tmp/token.txt)
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
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2026-01-19T12:30:00Z"
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
  }" | jq -r '.access_token' > /tmp/token.txt

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
  }"
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
  -H "Authorization: Bearer $TOKEN"
```

---

## Полный пример workflow

```bash
echo "=== Регистрация ==="
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"test123"}'

echo -e "\n\n=== Логин ==="
RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"test123"}')

echo $RESPONSE | jq .

TOKEN=$(echo $RESPONSE | jq -r '.access_token')
REFRESH_TOKEN=$(echo $RESPONSE | jq -r '.refresh_token')

echo -e "\n\n=== Получение текущего пользователя ==="
curl http://localhost:8080/api/v1/me \
  -H "Authorization: Bearer $TOKEN"

echo -e "\n\n=== Refresh токена ==="
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d "{\"refresh_token\":\"$REFRESH_TOKEN\"}"
```
