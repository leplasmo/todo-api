# ToDo API

## Usage

```bash

# set the json file to write to
export JSON_FILENAME=my-todos.json

# add a todo
go run cmd/todo/main.go -add "This is a new todo"

# or, you can add a todo by providing an interactive input
go run cmd/todo/main.go -add
This is another todo

# or, you can also pipe STDIN to add a new todo 
echo "This is yet another todo" | go run cmd/todo/main.go -add

# list todos
go run cmd/todo/main.go -list
#OUTPUT:
'
[ ]  1: This is a new todo
[ ]  2: This is another todo
[ ]  3: This is yet another todo
'
```
