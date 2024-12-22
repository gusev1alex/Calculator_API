# Calculator Web Service

## Описание
Этот проект предоставляет веб-сервис для вычисления арифметических выражений. Сервис принимает выражение через HTTP POST-запрос и возвращает результат вычисления или соответствующую ошибку.

## Стек технологий
- **Язык:** Go (Golang)
- **Фреймворк:** Gin

## Запуск проекта

### Установка
1. Убедитесь, что у вас установлен [Go](https://go.dev/dl/).
2. Клонируйте репозиторий:
   ```bash
   git clone <URL вашего репозитория>
   cd Calculator_API
   ```

### Запуск сервиса
Запустите сервер с помощью команды:
```bash
go run ./cmd/calc_service/main.go
```

Сервис будет доступен по адресу: `http://localhost:8080`

## API

### Endpoint
**POST** `/api/v1/calculate`

### Формат запроса
- **Headers**: `Content-Type: application/json`
- **Body**:
  ```json
  {
    "expression": "выражение для вычисления"
  }
  ```
  
### Примеры запросов

#### Успешное вычисление
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "3 + 5 * (2 - 8)"
}'
```
**Ответ:**
```json
{
  "result": "-13"
}
```

#### Ошибка 422 (некорректное выражение)
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "3 + 5 * (2 - a)"
}'
```
**Ответ:**
```json
{
  "error": "Expression is not valid"
}
```

#### Ошибка 500 (внутренняя ошибка сервера)
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": ""
}'
```
**Ответ:**
```json
{
  "error": "Internal server error"
}
```

## Тестирование
Для тестирования вы можете использовать Postman или `curl`. Примеры запросов приведены выше.

## Структура проекта
```
calc_service/
├── cmd/
│   └── calc_service/
│       └── main.go  # Точка входа
├── internal/
│   ├── calculator/
│   │   ├── calculator.go  # Логика вычисления
│   │   └── validator.go   # Валидация выражений
│   └── server/
│       ├── handler.go     # Обработчики HTTP-запросов
│       └── server.go      # Запуск сервера
├── go.mod                # Go-модуль
├── go.sum
└── README.md             # Документация
```

## Контакты
Если у вас есть вопросы или предложения, пожалуйста, свяжитесь со мной через GitHub Issues.

