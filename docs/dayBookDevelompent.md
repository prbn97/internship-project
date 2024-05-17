# A day book of this api development
- - -

Sumary


- - -
- - -




## April day 22 - What is a API?
Na última sexta recebi a tarefa de desenvolver uma API Rest. 

Recebi a recomendação de um amigo de usar OpenAPI pra planejar toda a API, 
com isso meu foco foi entender melhor o conceito e a história das API, API RestFul.

OpenAPI, resumidamente planeja toda uma api, alem de documenta-la. Isso produz uma aplicacao profissional. Entretanto o foco esta no principio de uma API.


## April day 24 - Getting API's examples
Busquei por *exemplos de api rest em Go*, na internet para criar a minha. Nao achei tanta coisa util devido ao uso de frameworks. Queria entender de forma mais conceitual
> devia ter buscando melhor, no decorrer dos dias, consegui.


## April day 27 of aprill - Get some milestones
Comecando a executar os primeiros itens do primeiro milestone definido. Criando a estrutura do todo item e buscando entrender melhor atraves de exemplos praticos como uma api funciona, como sua estrutura é organizada.

## April day 29  - Focus on develompment

# April day 30 - First Structure
The first structure that i follow.

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
            + /utils/                        
                - /http_errors.go
            - README.md
            - go.mod
            - Makefile

[Find this content](https://www.alura.com.br/conteudo/go-desenvolvendo-api-rest)
Helps me alot.

## May day 2 - Focus on development 

# May day 3 - Focus on *Milestone 1* and *Milestone 2* 

+ ## **fetaure - createDate and limitDate**
    - todo item have the date that the item was created, and the date limmit to be done.
    - if the is nil have no time limit.

+ ## **fetaure - Task list**
    - todo that is a big task, is a toDo item (with one more field ).

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
- - -

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
The todo_handler now are done. need 

## [need to focus on tests.](https://www.youtube.com/watch?v=xhgHeAhxizE)
### define tests http tests

i find good content in the YTB channel Golang Cafe.

[Go (Golang) Testing Tutorial](https://www.youtube.com/watch?v=LqU-0RVyq8I)

[Go (Golang) Test Coverage Tutorial](https://www.youtube.com/watch?v=xhgHeAhxizE)

[Go (Golang) httptest Tutorial](https://www.youtube.com/watch?v=LqU-0RVyq8I)


The new structure that i will follow.

    /todo-api/
        + /api/
                - /main.go
        + /docs/
                - /dayBookDevelopment.md
                - /documentation.md
        + /handlers/
                - /generateID.go
                - /generateID_test.go
                - /todoHandler.go
                - /todoHandler_test.go
                - /userHandler.go
                - /userHandler_test.go
        + /models/                        
                - /user.go
                - /todo.go
        + /tests/

        + /utils/                        
                - /http_errors.go
        - go.mod
        -god.sum
        - Makefile
        -README.md

# May day 7 - Focus on test development 

### Neet create unit tests to my heandler... and then create bigger test.
 Need check all milestones. Just then, add some feature.
___
 > Fetaure - createDate and limitDate

    - toDo item now have the date that has been created, and the date limmit to be finish.

    - if the field "date limmitis" is nil, have no time limit.
___
> Fetaure - Task list

    - Is a toDo item that have a list of toDo items.

    - Default task and Custom tasks.
___
> Fetaure - User can be created and can created tasks

[Go JWT Authentication Tutorial](https://www.youtube.com/watch?v=Qk6UgCps5Dc&t)

# May day 13 - Focus on test development

Problem with mi tests

this article can help me [link](https://golang.cafe/blog/golang-httptest-example.html)

# May day 14 - Focus on test development

TodoErro is handling errors in the requets, the goal is to get error with customize messages

i will start with the httpErrors.go, now the function recive a string, that will be for custom erros messages... like "id is not valid" or "id not found" for the api client.

```go

// function to handling errors with TodoError ***
func BadRequest(res http.ResponseWriter, req *http.Request, msg string) {
	res.WriteHeader(http.StatusBadRequest)
	errorJson := models.TodoError{ 
		Error:   "bad request",
		Message: msg,
	}
	jsonBytes, err := json.Marshal(errorJson)
	if err != nil {
		return
	}
	res.Write(jsonBytes)
}

func (h *TodoHandler) Create(res http.ResponseWriter, req *http.Request) {

	u := models.Todo{}
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		fmt.Println(err)
		utils.BadRequest(res, req, "invalid json") // customize the message ***
		return
	}
    
    .
    .
    .

}
```

# May day 15 - Focus on test development

Conversando com o Gui sobre minha api, vimos que posso deixar meu teste menores e mais claro, conversamos sobre set up enviroment.

Com isso relendo esse codigo vejo que "crio um manipulador" nos testes TestTodoHandler_Get e TestTodoHandler_Delete. Busquei mais informações com o que conversamos e entendi melhor, criar o ambiente pra teste e depois limpar os recursos, comumente chamado de "teardown"

## Setup Environment: 
A função setupEnv inicializa o TodoHandler e retorna uma função de teardown, mesmo que atualmente não haja necessidade de teardown. Isso prepara o ambiente de teste consistentemente para cada teste.

Vou ja criar um manipulador e adicionar itens nele melhorando os testes.

```go

// setupEnv initializes the environment for each test
func setupEnv(t *testing.T) (*TodoHandler, func()) {
	todoHandler := NewTodoHandler()

    //create todos
    //add todos in the handler

	// Return the handler and a teardown function to be called after the test
	return todoHandler, func() {
		// Perform any necessary teardown here
        // like delete the handler
	}
}

```