## Criacao de tipos

No go alem de ser fortemente tipado e ter tipos bem diferentes das linguagens comuns como float64, e possivel criar seus proprios tipos tambem.

### Type


```go
package main

type ID int

var f ID

func main() {
    f = 10
    println(f)
}
```