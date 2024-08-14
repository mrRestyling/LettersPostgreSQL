-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(16) NOT NULL,
    last_name VARCHAR(16) NOT NULL
);


CREATE TABLE IF NOT EXISTS letters (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    item VARCHAR(100),
    letter VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    mess VARCHAR(100),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
