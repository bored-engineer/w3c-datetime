# w3c-datetime [![Go Reference](https://pkg.go.dev/badge/github.com/bored-engineer/w3c-datetime.svg)](https://pkg.go.dev/github.com/bored-engineer/w3c-datetime)
Implements the [W3C datetime format](https://www.w3.org/TR/NOTE-datetime) for Golang.

```go
package main

import (
	"fmt"

	datetime "github.com/bored-engineer/w3c-datetime"
)

func main() {
	d, err := datetime.Parse("1994-11-05")
	if err != nil {
		panic(err)
	}
	fmt.Println(d.Precision) // YYYY-MM-DD
    fmt.Println(d.Time) // 1994-11-05 00:00:00 +0000 UTC
 
	dt, err := datetime.Parse("1994-11-05T08:15:30Z")
	if err != nil {
		panic(err)
	}
	fmt.Println(dt.Precision) // YYYY-MM-DDThh:mm:ssTZD
    fmt.Println(dt.Time) // 1994-11-05 08:15:30 +0000 UTC
}
```
