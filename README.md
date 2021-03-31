## TODO
- [ ] Variable Arguments in Function, Variable Parameters in Function Call (func Test(args ...int) , Test(args[:]...)) 
```go
package main

import (
  "fmt"
)

func Test(args ...string) {
   for _, v := range args {
      fmt.Println(v)
   }
}

func main() {
   Test([]string{"Hello", "World", "!"}...)
}
```
