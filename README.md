# Todo List in Go

A simple Todo List implemented in Go. This project provides a basic Todo List functionality implemented in Go. It allows users to manage their todo items by performing CRUD (Create, Read, Update, Delete) operations via HTTP endpoints.

## Usage

To run the Todo List server, execute the following command:

```
go run cmd/todo-api/main.go
```

## API Endpoints

- **GET /todos**: Retrieve all todo items.
- **POST /todos/add**: Add a new todo item.
- **PUT /todos/update**: Update an existing todo item.
- **DELETE /todos/delete**: Delete an existing todo item.

## Using CLI (curl)

- Get all todo items:
  ```
  curl http://localhost:8080/todos
  ```

- Add a new todo item:
  ```
  curl -X POST -H "Content-Type: application/json" -d '{"id": "1", "title": "New Todo", "done": false}' http://localhost:8080/todos/add
  ```

- Update a todo item:
  ```
  curl -X PUT -H "Content-Type: application/json" -d '{"id": "1", "title": "Updated Todo", "done": true}' http://localhost:8080/todos/update
  ```

- Delete a todo item:
  ```
  curl -X DELETE -H "Content-Type: application/json" -d '{"id": "1"}' http://localhost:8080/todos/delete
  ```

## Using Postman

- Get all todo items:
    - Method: GET
    - URL: http://localhost:8080/todos


- Add a new todo item:
    - Method: POST
    - URL: http://localhost:8080/todos/add
    - Headers:
        - Key: Content-Type
        - Value: application/json
    - Body (raw JSON):
      ```
      {
          "id": "1",
          "title": "New Todo",
          "done": false
      }
      ```

- Update a todo item:
    - Method: PUT
    - URL: http://localhost:8080/todos/update
    - Headers:
        - Key: Content-Type
        - Value: application/json
    - Body (raw JSON):
      ```
      {
          "id": "1",
          "title": "Updated Todo",
          "done": true
      }
      ```

- Delete a todo item:
    - Method: DELETE
    - URL: http://localhost:8080/todos/delete
    - Headers:
        - Key: Content-Type
        - Value: application/json
    - Body (raw JSON):
      ```
      {
          "id": "1"
      }
      ```


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

by **Eduardo Raider** - Software Engineer