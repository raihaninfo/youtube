## REST API with mux

This is a simple REST API with [mux](https://github.com/gorilla/mux).

## api endpoints table

|     | method | endpoint       | description       |
| --- | ------ | -------------- | ----------------- |
| 1   | GET    | /api/user      | get all users     |
| 2   | GET    | /api/user/{id} | get user by id    |
| 3   | POST   | /user          | create new user   |
| 4   | PUT    | /api/user/{id} | update user by id |
| 5   | DELETE | /api/user/{id} | delete user by id |

## Api structure

```go
type User struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Email  string `json:"email"`
    Password   string `json:"password"`

}
```
