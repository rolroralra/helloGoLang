## TODO
- [ ] Variable Arguments in Function, Variable Parameteres in Function Call (func Test(args ...int) , Test(args[:]...)) 
```go
pacakge main

import (
  "fmt"
)

func Test(args ...string) {
   for _, v := range args {
      fmt.Println(v)
   }
}

func main() {
   test([]string{"Hello", "World", "!"}...)
}
```
