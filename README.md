# Dynamic user segmentation service
This is test task for avito tech.
[x] Main task
[x] Implement saving the user's hit/drop history from the segment
[x] Implement the ability to set TTL
[x] Automatically adding a user to a segment
[ ] Testing
[ ] Swagger documentation

# Get started
To start application

```
docker-compose up
```
Fill database
```
docker exec -i dynamic_user_segmentation_service-db-1 psql -U postgres < migrations/000001_init.up.sql
```
## Requests
### POST http://localhost:8080/users/
```
{
    "id":1
}
```
Response:
```
{
    "id":1
}
```
Error response:
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
Response:
```
{
    "id":3,
    "title":"qwerty2"
}

```
Error response:
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
Response:
```
{
    "status":"done"
}

```
Error response:
```
{
    "error":"not found"
}
```
### PUT http://localhost:8080/slugsUsers/
title_add - array of titles to add
title_delete - array of titles to delete
id - users id
ttl - time to live in seconds

```
{
    "title_add": ["qwerty1"],
    "title_delete": [],
    "id": 1,
    "ttl": 0
}
```
Response:
```
{
    "status":"done"
}

```
Error response:
```
{
    "error":"already exists"
}
```
Request:
```
{
    "title_add": [],
    "title_delete": ["qwerty1"],
    "id": 1,
    "ttl": 0
}
```
Response:
```
{
    "status":"done"
}

```
Error response:
```
{
    "error":"not found"
}
```
Request:
```
{
    "title_add": ["qwerty1"],
    "title_delete": [],
    "id": 1,
    "ttl": 5
}
```
Response:
```
{
    "status":"done"
}

```
Error response:
```
{
    "error":"already exists"
}
```
5 seconds after adding the log to the console:
```
time="2023-08-31T18:42:44Z" level=info msg=deleted
```
### GET http://localhost:8080/slugsUsers/
```
{
    "id": 1
}
```
Response:
```
{
    "user_id":1,
    "slugs":["qwerty3"]
}
```
Error response if user or users slugs does not exist:
```
{
    "error":"not found"
}
```
### GET http://localhost:8080/records/
{
    "date": "2023-08"
}
Response:
Link in browser:
![link screenshot](https://github.com/neglarken/dynamic_user_segmentation_service/raw/master/link_screen.png)
### GET http://localhost:8080/files/
Response:
Link in browser:
![link screenshot](https://github.com/neglarken/dynamic_user_segmentation_service/raw/master/link_screen.png)