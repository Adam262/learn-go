package main

import (
  "fmt"

  "github.com/user/hello/arrays"
  "github.com/user/hello/loops"
  "github.com/user/hello/mathstuff"
  "github.com/user/hello/morestrings"
  "github.com/user/hello/pointers"
  "github.com/user/hello/structs"
)

func main() {
  // defer will not run until surrounding function returns
  // they can be stacked and called in LIFO order
  defer fmt.Println("Deferred World")
  defer fmt.Println("Deferred Hello")

  arrays.Main()

  loops.Main()

  mathstuff.Main()

  morestrings.Main()

  pointers.Main()

  structs.Main()
}
