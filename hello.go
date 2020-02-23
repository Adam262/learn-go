package main

import (
  "fmt"

  "github.com/user/hello/concurrency"
)

func main() {
  // defer will not run until surrounding function returns
  // they can be stacked and called in LIFO order
  defer fmt.Println("Deferred World")
  defer fmt.Println("Deferred Hello")

  // arrays.Main()

  concurrency.Main()

  // funcs.Main()

  // interfaces.Main()

  // loops.Main()

  // maps.Main()

  // mathlocal.Main()

  // methods.Main()

  // pointers.Main()

  // regexplocal.Main()

  // stringslocal.Main()

  // structs.Main()
}

// Concurrence - channels
