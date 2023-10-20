## Maps

```go
package main

import "fmt"

func main() {
	salarios := map[string]int{"Gabriel": 10, "Thais": 10}
	fmt.Println(salarios["Gabriel"])
	delete(salarios, "Gabriel")
	salarios["Novo"] = 20
	fmt.Println(salarios["Novo"])

	for nome, salario := range salarios {
		fmt.Printf("O nome e: %s e o salario e: %d\n", nome, salario)
	}
}

```