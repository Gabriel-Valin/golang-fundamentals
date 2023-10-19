## Percorrendo arrays

```go
package main

import (
    "fmt"
)

func main() {
    // array tem valor finito
    var meuArray [3]int
    meuArray[0] = 20
    meuArray[1] = 40
    meuArray[2] = 60

    for i, v := range meuArray {
        fmt.Printf("O valor do indice eh %d e o valor eh %d\n", i, v)
    }
}