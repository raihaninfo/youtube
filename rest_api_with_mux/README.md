## REST API with mux

This is a simple REST API with [mux](https://github.com/gorilla/mux).



## api endpoints table

|     | method | endpoint        | description       |
| --- | ------ | --------------- | ----------------- |
| 1   | GET    | /api/books      | get all books     |
| 2   | GET    | /api/books/{id} | get book by id    |
| 3   | POST   | /book           | create new book   |
| 4   | PUT    | /api/book/{id}  | update book by id |
| 5   | DELETE | /api/book/{id}  | delete book by id |


## Api structure

```go
type Book struct {
    ID      string `json:"id"`
    Name    string `json:"title"`
    Author  string `json:"author"`
    Price   string `json:"price"`
    Status  string `json:"status"`
}
```
