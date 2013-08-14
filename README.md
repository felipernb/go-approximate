go-approximate
==============

Small Go library to format numbers approximately in a "human readable" way ("7.5K", "280M", etc.)

Based on https://github.com/mahemoff/approximate

Usage
-----

```
    go get github.com/felipernb/go-approximate
```

myfile.go

```go
import (
    "fmt"
    "github.com/felipernb/go-approximate"
)

func main() {
    fmt.Println(approximate.Approximate(949384902)) // 9.5G
}
```
