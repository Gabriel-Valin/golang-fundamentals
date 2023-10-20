## Funcoes

### Funcao padrao

```go
package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
}

func sum(a, b int) int {
	return a + b
}
```

### Funcao com dois retornos

```go
package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
}

func sum(a, b int) (int, bool) {
    if a+b >= 10 {
        return a+b, true
    }
	return a + b, false
}
```

### Funcao retornando erro (no Go nao tem exception, trycatch)

```go
package main

import "fmt"

func main() {
	valor, err := sum(10, 0)
	valor2, err2 := sum(8, 1)

	if err != nil {
		fmt.Println(err)
	}

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(valor)
	fmt.Println(valor2)
}

func sum(a, b int) (int, error) {
	if a+b >= 10 {
		return 0, errors.New("O valor e maior que 10")
	}
	return a + b, nil
}
```

### Funcao com n parametros

```go
package main

import "fmt"

func main() {
    fmt.Println(sum(13,25,31,3251321,4213,213,1))
}

func sum(numeros ...int) int {
    total := 0
    for _, numero := range numeros {
        total += numero
    }

    return total
}
```