package main

import (
  "fmt"

  "github.com/user/hello/maps"
)

func main() {
  // defer will not run until surrounding function returns
  // they can be stacked and called in LIFO order
  defer fmt.Println("Deferred World")
  defer fmt.Println("Deferred Hello")

  // arrays.Main()

  // loops.Main()

  maps.Main()

  // mathstuff.Main()

  // morestrings.Main()

  // pointers.Main()

  // structs.Main()
}
