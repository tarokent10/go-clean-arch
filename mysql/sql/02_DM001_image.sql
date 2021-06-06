use mydb;
-- user
INSERT INTO user (
    user_id, name, password, created_at, updated_at
) VALUES (
    'test', "taro tanaka", "1234", now(), now()
)

