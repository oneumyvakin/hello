# hello
Hello World example on Go

## Installing

Install in the usual Go way:

```sh
go get -u github.com/oneumyvakin/hello
```

## Examples

Just "Hello, World!" in English:
```go
package main

import (
    "fmt"
    "github.com/oneumyvakin/hello"
)	

func main() {
    err := hello.World()
    if err != nil {		
        fmt.Printf("failed hello world: '%s'", err)
        return
    }
}
```

Use one of supported language:
```go
package main

import (
    "fmt"
    "github.com/oneumyvakin/hello"
)	

func main() {
    myLang := "ru-RU"
    err := hello.WorldIn(myLang)
    if err != nil {		
        fmt.Printf("failed hello world with lang '%s': '%s'", myLang, err)
        return
    }
}
```