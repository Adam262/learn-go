package main

import (
  "fmt"

  "github.com/user/hello/arrays"
  "github.com/user/hello/funcs"
  "github.com/user/hello/loops"
  "github.com/user/hello/maps"
  "github.com/user/hello/mathlocal"
  "github.com/user/hello/pointers"
  "github.com/user/hello/stringslocal"
  "github.com/user/hello/structs"
)

func main() {
  // defer will not run until surrounding function returns
  // they can be stacked and called in LIFO order
  defer fmt.Println("Deferred World")
  defer fmt.Println("Deferred Hello")

  arrays.Main()

  funcs.Main()

  loops.Main()

  maps.Main()

  mathlocal.Main()

  stringslocal.Main()

  pointers.Main()

  structs.Main()
}
