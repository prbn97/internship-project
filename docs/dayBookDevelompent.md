# A day book of this api development
- - -

Sumary


- - -
- - -




## April day 22 - What is a API?
Last Friday I received the task of developing a Rest API.

We received a recommendation from a friend to use OpenAPI to plan the entire API,
Therefore, my focus was to better understand the concept and history of the API, RestFul API.

OpenAPI, in short, plans an entire API, in addition to the documentation. This produces a professional application. However, the focus is not on the principle of an API.


## April day 24 - Getting API's examples
I searched for *examples of api rest in Go* on the internet to create my own. I didn't find much useful due to the use of frameworks. I would like to understand in a more conceptual way
> I should have tried better, as the days went by, I did it.


## April day 27 of aprill - Get some milestones
Starting to execute the first items of the first defined milestone. Creating the structure of the entire item and seeking to better understand through practical examples how an API works, how its structure is organized.

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

Talking to Gui about my API, we saw that I can make my test smaller and clearer, we talked about setting up environment.

So, re-reading this code I see that I "create a handler" in the TestTodoHandler_Get and TestTodoHandler_Delete tests. I sought more information from what we talked about and understood better, creating the environment for testing and then cleaning up the resources, commonly called "teardown"

## Setup Environment: 
The setupEnv function initializes the TodoHandler and returns a teardown function, even if there is currently no need for a teardown. This prepares the test environment consistently for each test.

I'm going to create a handler and add items to it, improving the tests.

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


# May day 20 - **Focus Milestone 4: Error Handling**

Today my focus in Error Handling reval that i have a problem with my PUT method, need fix that.

Just then im getting back to handling bad request like using invalid ids when doing get by id.

# May day 21 - **Focus Milestone 4: Error Handling**

I found a problem in my PUT and POST methods. When I make a request and put one of the fields wrong, it doesn't show any error, it just ignores that wrong field...

> example to make it clearer.
```shell
    curl --request POST \
    --data '{"titl7e":"title", "description": "description example"}' \
    http://localhost:8080/todos
```

it will create a toDO with an empty *"title"*... ​​and ignore the wrong field that I entered.

The solution to this was to create a decoder and use the DisallowUnknownFields() function to check the request body
```go
decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
```