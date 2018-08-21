# str
The `str` package contains some mid-level string operations.

## Imports
```go
import "github.com/suite911/str911/str"
```

## Usage Examples
```go
package main

import (
	"fmt"

	"github.com/suite911/str911/str"
)

func main() {
	match, tail := str.CaseTrimPrefix("Hello, World!", "heLLo")
	simp := str.Simp("@Amy.Adzuki#1234!")
	fmt.Printf("%v: \"%s\"; \"%s\"\n", match, tail, simp)
	// Prints: `true: ", World!"; "amy.adzuki1234"`
}
```
