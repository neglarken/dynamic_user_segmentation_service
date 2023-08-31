# Dynamic user segmentation service
Это мое решение тестового задания от Avito Tech
- [x] Основное задание
- [x] Реализовать сохранение истории посещений/удалений пользователя из сегмента
- [x] Реализовать возможность установки TTL
- [x] Автоматическое добавление пользователя в сегмент
- [ ] Тесты
- [ ] Swagger документация

# Get started
Запустить приложение

```
docker-compose up
```
Заполнить базу данных
```
docker exec -i dynamic_user_segmentation_service-db-1 psql -U postgres < migrations/000001_init.up.sql
```
## Заросы
### POST http://localhost:8080/users/
```
{
    "id":1
}
```
Ответ:
```
{
    "id":1
}
```
Ответ ошибки:
```
{
    "error":"already exists"
}
```
### POST http://localhost:8080/slugs/
```
{
    "title": "qwerty1",
    "part": 0
}
```
Ответ:
```
{
    "id":3,
    "title":"qwerty2"
}

```
Ответ ошибки:
```
{
    "error":"already exists"
}
```
### DELETE http://localhost:8080/slugs/
```
{
    "title": "qwerty2"
}
```
Ответ:
```
{
    "status":"done"
}

```
Ответ ошибки:
```
{
    "error":"not found"
}
```
### PUT http://localhost:8080/slugsUsers/
- title_add - массив названий сегментов для добавления
- title_delete - массив названий сегментов для удаления
- id - id пользователя
- ttl - время жизни в секундах

```
{
    "title_add": ["qwerty1"],
    "title_delete": [],
    "id": 1,
    "ttl": 0
}
```
Ответ:
```
{
    "status":"done"
}

```
Ответ ошибки:
```
{
    "error":"already exists"
}
```
Запрос:
```
{
    "title_add": [],
    "title_delete": ["qwerty1"],
    "id": 1,
    "ttl": 0
}
```
Ответ:
```
{
    "status":"done"
}

```
Ответ ошибки:
```
{
    "error":"not found"
}
```
Запрос:
```
{
    "title_add": ["qwerty1"],
    "title_delete": [],
    "id": 1,
    "ttl": 5
}
```
Ответ:
```
{
    "status":"done"
}

```
Ответ ошибки:
```
{
    "error":"already exists"
}
```
Прошло 5 секунд после выполнения запроса:
```
time="2023-08-31T18:42:44Z" level=info msg=deleted
```
### GET http://localhost:8080/slugsUsers/
```
{
    "id": 1
}
```
Ответ:
```
{
    "user_id":1,
    "slugs":["qwerty3"]
}
```
Ответ если пользователь не существует или он не привязан ни к одному сегменту:
```
{
    "error":"not found"
}
```
### GET http://localhost:8080/records/
{
    "date": "2023-08"
}
Ответ:
Ссылка на странице браузера:
![link screenshot](https://github.com/neglarken/dynamic_user_segmentation_service/blob/main/link_screen.png)
### GET http://localhost:8080/files/
Ответ - ссылка на странице браузера:
![link screenshot](https://github.com/neglarken/dynamic_user_segmentation_service/blob/main/link_screen.png)

### Вопросы, возникшие в ходе решения задачи
- В каком виде возвращать ссылку на csv файл?
    - Был вариант отдавать ссылку на файл в теле ответа, но тогда просмотр содержимого файла осуществлялся бы в окне браузера.
Но я от него отказался, так как посчитал переадресацию на страницу со ссылкой, по которой можно скачать файл, более удобным способом.
- Как удобно реализовать заполнение базы данных таблицами из sql файла?
    - Долгий поиск в интернете дал множество плодов, но ни один из них не работал у меня, поэтому было решено заполнять инициализированную
базу командой миграции, несмотря на то, что миграции не требовались в ТЗ.
- Что должно выполнять механизм TTL? База данных или приложение?
    - Написание триггеров в базе данных оказалось слишком долгим, поэтому было решено переложить ответственность за реализацию механизма TTL
на приложение.
