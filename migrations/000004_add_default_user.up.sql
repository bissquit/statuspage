-- Создаём тестового пользователя для документации и разработки
-- Email: user@example.com
-- Password: user123
-- ВАЖНО: Только для development/testing!

INSERT INTO users (email, password_hash, first_name, last_name, role)
VALUES (
           'user@example.com',
           '$2a$10$G7I/uOGevrmM6Q0QTyhnQerwJbSRDQyti97v7jpLbsFLAN5CXOCnO',
           'Test',
           'User',
           'user'
) ON CONFLICT (email) DO NOTHING;
