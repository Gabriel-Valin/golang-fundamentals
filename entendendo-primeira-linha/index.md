### Pacotes

```go
package main

func main() {
    println("Hello World!")
}
```

Este codigo acima se refere a um arquivo chamado `main.go` que todo projeto go tem que ter o arquivo "index.js" do projeto. Ele sempre tera esse package main em sua primeira linha.

### Packages por pastas

Se tivermos a seguinte estrutura de diretorios `projeto/services/get-user.go` o package que o arquivo `get-user.go` devera ter e o *service*, pois e a pasta em que o arquivo se encontra.

`get-user.go`

```go
package services

func getUser() {
    println("User")
}
```