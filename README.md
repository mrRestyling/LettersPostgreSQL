# LettersPostgreSQL.

Этот проект предназначен для демонстрации использования PostgreSQL с языком GO.

## Используемые технологии
-  **PostgreSQL** - мощная и надежная система управления базами данных, поддерживающая SQL и предоставляющая широкие возможности для хранения и обработки данных.
- **Фреймворк Echo**
- **Архитектурные слои** - хендлеры, сервисный уровень, подключение к БД, модели
- **Graceful shutdown**
- **Docker**
- **Postman**


## Функционал

Проект представляет собой веб-сервер, построенный с использованием фреймворка Echo на языке Go. 
- Обработка имен
- Обработка писем
- Обработка сообщений





## Пример использования

*Docker*
1) Подключение к базе данных через Docker:
docker-compose up -d
2) Отключение:
docker-compose down (Дополнение: флаг для удаления данных: -v)

*Postman*

1) Тестовый POST-запрос на http://localhost:8080/name:
{
"first": "Иван",
"last": "Иванов"
}

Ответ сервера: 1 (уникальный id в БД)

2) Тестовый GET-запрос на http://localhost:8080/letter:
{
  "user_id": 1
}


Ответ сервера:

{
    "Amount": 3,
    "Letters": [
        
        "Письмо № 1 || Заголовок: Маме || Содержание письма: Мы на море,у нас все хорошо. Мы вам привезем клубнику",
        "Письмо № 2 || Заголовок: Соседу || Содержание письма: Перестаньте шуметь после 23:00",
        "Письмо № 3 || Заголовок: Ирине || Содержание письма: Ваши письма в красивом формате!"  
  ]
}


