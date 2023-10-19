## Declaracao e atribuicao

### Declaracao simples
```go
package main


// Declaracao
// b vai ser sempre false caso nao seja atribuido nenhum valor a ela, o go infere o valor na variavel.
var b bool

// Constante nunca pode ter seu valor alterado, nao e permitido.
const a = "Hello world!"

func main() {
    println(b)
}
```


### Declaracao multipla
```go
package main

var (
    b bool // sempre inferido (false)
    c int // sempre inferido (0)
    d string // sempre inferido ("")
    e float64 // sempre inferido (0.000000+)

    // todas essas vars sao de escopo global
)

func main() {
    println(d)
}
```


### Declaracao multipla
```go
package main

var (
    b bool // sempre inferido (false)
    c int // sempre inferido (0)
    d string // sempre inferido ("")
    e float64 // sempre inferido (0.000000+)

    // todas essas vars sao de escopo global
)

func main() {
    println(d)
}
```


### Short hand :=
```go
package main

func main() {
    d := "Hello world!" // o go ja faz inferencia do tipo string
    d := 2 // isso nao funcionara, so pode se usar essa syntax na primeira declaracao.
    d = "Mamacita!" // e o jeito correto de usar.
    
    println(d)
}
```