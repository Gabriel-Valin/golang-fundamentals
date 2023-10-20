## Ponteiros

```go
package main
 
import "fmt"
 
func main() {
    // Memoria -> endereco -> valor
    x := 10
    var ponteiro *int = &x
    println(ponteiro)
    *ponteiro = 20
    b := &x // b = endereco da variavel x na memoria, ponteiro
    *b = 30 // o valor da variavel que possui o endereco vai ser 30
    println(x)
}
```

### Outro exemplo

```go
package main
 
import "fmt"
 
func soma(a, b *int) int {
    *a = 50
    *b = 10

    return *a + *b
}

func main() {
    minhaVar := 20
    minhaVar2 := 30
    soma(&minhaVar, &minhaVar2)

    println(minhaVar)
    println(minhaVar2)
}
```