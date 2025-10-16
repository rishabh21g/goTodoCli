# goTodoCli

A simple terminal-based TODO CLI written in Go.

## Screenshot
![TODO CLI Application](docs/todo-cli.png)

## Build and Run (Go)
```bash
go build -o app .
./app -list
```

## Run directly
```bash
go run . -list
```

## Docker
```bash
docker build -t go-todo-cli .
docker run --rm -it go-todo-cli -list
```

## Features
- List tasks
- Add tasks
- Mark complete
- Delete tasks

## Notes
- Place the screenshot at: docs/todo-cli.png
- Requires Go and modules enabled.