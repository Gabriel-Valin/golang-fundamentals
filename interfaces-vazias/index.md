## Interfaces vazias

```go
package main

import "fmt"

func main() {
    var x interface{} = 10
    var y interface{} = "Hello world!"

    showType(x)
    showType(y)
}

func showType(t interface{}) {
    fmt.Println("o tipo e %T e o valor e %v", t, t)
}
```