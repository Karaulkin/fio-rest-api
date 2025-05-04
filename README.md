# FIO Service

Сервис ФИО с REST API при помощи *Echo* для управления пользователями и обогощения данными.

## Требования

- Go 1.2+
- PostgresSQL 13+

## Установка и запуск

1. Клонируйте репозиторий:

>если запускаете локально в .env CONFIG_PATH будет  CONFIG_PATH=./config/local.yaml

2. Установите зависимости:
   ```bash
    go mod tidy
   ```

3. Запустите сервер:
   ```bash
    go run cmd/server/main.go
   ```

Сервер будет доступен по адресу: http://localhost:8080/api/v1

## API документация

Swagger документация доступна по адресу: http://localhost:8080/swagger/index.html

## Основные функции

1. **Управление пользователями**
    - Для получения данных с различными фильтрами и пагинацией
    - Для удаления по идентификатору 
    - Для изменения сущности
    - Для добавления новых людей в формате
    ```json
    {
        "name": "Dmitriy",
        "surname": "Ushakov",
        "patronymic": "Vasilevich" // необязательно
    }
    ```
## Запуск при помощи Docker

При запуске в контейнере поменять в .env

`CONFIG_PATH=./config/dev.yaml`

   ```shell
      docker compose up
   ```
