## Structs

Basicamente um agrupamento de dados

```go
package main
 
import "fmt"
 
type endereco struct {
    rua string
    numero init
    cidade string
}
 
func main() {
    endereco := endereco{
        rua: "Rua tal",
        numero: 10,
        cidade: "Botucatu",
    }
 
    fmt.Println(endereco)
    endereco.numero = 18
    fmt.Println(endereco.numero)
}
```