# A day book of this api development


### 22 de abril

Na última sexta recebi a tarefa de desenvolver uma api. Recebi a recomendação de um amigo de usar OpenAPI pra planejar toda a API, com isso meu foco foi entender melhor o conceito e a história das API, API RestFul e como usar a OpenAPI, achei conteúdos interessantes.
[video 1](https://www.youtube.com/watch?v=umaXYEbd5vA)
[video 2](https://www.youtube.com/watch?v=lsMQRaeKNDk&list=PLOspHqNVtKAAAq9pHWlEiRUVcYMCcu4X0&index=4)
[video 3](https://www.youtube.com/watch?v=pRS9LRBgjYg&list=PLOspHqNVtKAAAq9pHWlEiRUVcYMCcu4X0&index=5)

***Revisando sobre o texto, os planos com o decorrer do projeto mudaram. A ideia é entender principios, então desenvolver a api sem recursos foi o caminho seguido***
### 24 de abril

Buscando [exemplos](https://www.youtube.com/watch?v=7cTpMZdZpzM) de api rest em Go na internet para criar a minha.


### 27 de abril
Comecando a executar os primeiros itens do primeiro milestone definido. Criando a estrutura do todo item e buscando entrender melhor atraves de exemplos praticos como uma api funciona, como sua estrutura é organizada.

### 29 de abril
Focus on develompment today. Hoje inicie a criacao do metodo POST

### 30 de abril
one structure that i will follow (and update).

    todo-api/
    - /api/
        - /main.go
        - /todo_handler.go
    - /models/
        - /todo.go
    - /tests/
        - /todo_handler_test.go
    - README.md
    - go.mod
    - Makefile

[Find this content](https://www.alura.com.br/conteudo/go-desenvolvendo-api-rest)
Helps me alot.

### 2 de abril

lets finish this week with *Milestone 1* and *Milestone 2* finished.

### 3 de abril

#### fetaure - Task list
todo that is a list of todo itens, have all proprieties.

#### fetaure - Date and Date limit
todo item have the date that the item was created, and the date limmit, if the is nil dont have time limit

to get a better code i need improve my http requests... get some knowledge in this topic.

+ My methods POST and GET need more error handling
+ My *todo list* need better code
    - my api can have many list of *todo list*