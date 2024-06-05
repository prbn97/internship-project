# Task list API

Documentation goes here


POST tasks

```shell
curl -v POST \
  -d '{"title":"Task 1", "description": "description 1"}' \
  http://localhost:8080/tasks


curl -v POST \
  -d '{"title":"title", "description": ""}' \
  http://localhost:8080/tasks
```

GET tasks

```shell
curl -v GET http://localhost:8080/tasks


curl -v GET \
  http://localhost:8080/todos/{id}
```

PUT tasks

```shell
curl -v PUT \
  http://localhost:8080/todos/{id}/complete

curl -v PUT \
  -d '{"title":"Title Update", "description":"Description Update"}' \
  http://localhost:8080/tasks/{id}

curl -v PUT \
  -d '{"title":""}' \
  http://localhost:8080/tasks/{id}


curl -v PUT \
  -d '{"description":""}' \
  http://localhost:8080/tasks/{id}
```

DELETE tasks

```shell
curl -v DELETE \
  http://localhost:8080/tasks/{id}

```
curl -v -X PUT \
  -d '{"title":""}' \
  http://localhost:8080/tasks/89d9777c857a7fc95844