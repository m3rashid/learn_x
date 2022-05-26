```
go get github.com/99designs/gqlgen

printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
go mod tidy

go run github.com/99designs/gqlgen init

go run server.go
```
