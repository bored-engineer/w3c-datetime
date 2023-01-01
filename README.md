# w3c-datetime
Implements the [W3C datetime format](https://www.w3.org/TR/NOTE-datetime) for Golang.

```go
package main

import (
	"fmt"

	datetime "github.com/bored-engineer/w3c-datetime"
)

func main() {
	low, err := datetime.Parse("1994-11-05")
	if err != nil {
		panic(err)
	}
	fmt.Println(low.Precision, low.Time)
	high, err := datetime.Parse("1994-11-05T08:15:30Z")
	if err != nil {
		panic(err)
	}
	fmt.Println(high.Precision, high.Time)
}
```
