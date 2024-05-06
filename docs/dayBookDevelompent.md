# A day book of this api development
- - -

Sumary


- - -
- - -




### April day 22 - What is a API?
- - -
Na última sexta recebi a tarefa de desenvolver uma API Rest. 

Recebi a recomendação de um amigo de usar OpenAPI pra planejar toda a API, 
com isso meu foco foi entender melhor o conceito e a história das API, API RestFul.

OpenAPI, resumidamente planeja toda uma api, alem de documenta-la. Isso produz uma aplicacao profissional. Entretanto o foco esta no principio de uma API.


### April day 24 - Getting API's examples 
- - -
Busquei por *exemplos de api rest em Go*, na internet para criar a minha. Nao achei tanta coisa util devido ao uso de frameworks
> devia ter buscando melhor, dias depois consegui


### April day 27 of aprill - Get some milestones
- - -
Comecando a executar os primeiros itens do primeiro milestone definido. Criando a estrutura do todo item e buscando entrender melhor atraves de exemplos praticos como uma api funciona, como sua estrutura é organizada.

### April day 29  - Focus on develompment
- - -
### April day 30 - First Structure
- - -
The structure that i will follow.

    /todo-api/
            + /api/
                - /main.go
            + /docs/
                - /dayBookDevelopment.md
                - /documentation.md
            + /handlers/
                - /user_handler.go
                - /todo_handler.go
            + /models/                        
                - /user.go
                - /todo.go
            + /tests/
            + /utils/                        
                - /http_errors.go
            - README.md
            - go.mod
            - Makefile

[Find this content](https://www.alura.com.br/conteudo/go-desenvolvendo-api-rest)
Helps me alot.

- - -
### May day 2 - Focus on development 
- - -
### May day 3 - Focus on *Milestone 1* and *Milestone 2* 
- - -
+ #### **(future) fetaure - Task list**
    - todo that is a list of todo itens, have all proprieties.

+ #### **(future) fetaure - Date and Date limit**
    - todo item have the date that the item was created, and the date limmit.
    - if the is nil have no time limit.
_ _ _
> Problem

####  Get good http request methods. I was not have a good use of the standard lib of Go... get some knowledge in this topic.
+ My methods POST and GET need more error handling
+ My *todo list* need better code
    - my api can have many list of *todo list*
- - -
> Search
### find this article [An Introduction to Handlers and Servemuxes in Go](https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go) too

    "Processing HTTP requests with Go is primarily about two things: handlers and servemuxes."

    Whereas a servemux (also known as a router) stores a mapping between the predefined URL paths for your application and the corresponding handlers. Usually you have one servemux for your application containing all your routes.

    Go's net/http package ships with the simple but effective http.ServeMux servemux, plus a few functions to generate common handlers including http.FileServer(), http.NotFoundHandler() and http.RedirectHandler().
    
    Let's take a look at a simple (but slightly contrived!) example which uses these:


> Solution

 ### Find this [video](https://www.youtube.com/watch?v=8S30eoBSojU) - how make a api without frameworks. 

 > Is from a playlist - [*Master the Go http packages*](https://www.youtube.com/watch?v=u3YWN4TF81w&list=PLLf6iaZKV_xuD2D-7UkK_ToRwBBc8nv9P&index=1).

- - - 
## http.ServMux() + type Handler interface { }
ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

    package http // import "net/http"

        type Handler interface {
            ServeHTTP(ResponseWriter, *Request)
        }
    
    

### A http.ServMux() is a http serv that associates requests to respond to *Handler* INTERFACE.

### *Handler* are our requests methods, that will be our Fucntions to deal with requests. When certan requests are passes in our server lets use this way to deal with it.
- - -
### Start a http.ServMux() 


# May day 5 - Focus on development 

the videdo make me do a user_handler that can be use for and user of the ToDO items JSON [Web Tokens](https://www.alura.com.br/artigos/o-que-e-json-web-tokensg)


# May day 6 - Focus on development 
need create the todo_handler now and focus on tests