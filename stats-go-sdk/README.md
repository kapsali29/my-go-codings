# Golang Stats SDK

```shell
go run .\main.go
```

# Dependency Management in Go projects

`go get` adds or changes dependency versions, while `go build` and `go test` can add needed dependencies automatically.` go mod tidy` is the cleanup step that adds missing requirements and removes unused ones, and it’s good practice to run it before committing

```shell
go mod init your/module
go get example.com/some/dependency@latest
go mod tidy
go build ./...
go test ./...
```

# Execution
```shell
go run .\main.go answers -from "2026-03-05 10:00:00" -to "2026-03-07 10:00:00"

go run .\main.go questions -list
go run .\main.go questions -id 79907160 -display
go run .\main.go questions -id 79907160 -answers
```