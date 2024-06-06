# Task list API

Documentation goes here



## Request Examples

### POST tasks

```shell
curl -v -X POST -d '{"title":"Title"}' \
  http://localhost:8080/tasks


curl -X POST -d '{"title":"Task 2", "description": "description of the task 2"}' \
  http://localhost:8080/tasks
```

### GET tasks

```shell
curl -v -X GET http://localhost:8080/tasks
```


### PUT tasks

```shell
curl -X PUT -d '{"title":""}' \
  http://localhost:8080/tasks/{id}


curl -X PUT -d '{"description":""}' \
  http://localhost:8080/tasks/{id}


curl -X PUT http://localhost:8080/tasks/{id}
```

### DELETE tasks

```shell
curl -X DELETE \
  http://localhost:8080/tasks/{id}

```
