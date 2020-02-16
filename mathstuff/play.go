package mathstuff

import (
  "fmt"

  "github.com/user/hello/loops"
)

func Main() {
  // var keyword variable declaration with explicit type + initial value
  var x int = 2

  // var keyword variable declaration with explicit type but no initial value
  var y int
  y = 5

  // short form variable declaration with implicit type.
  // short form is only available within a function
  sum := Add(x, y)
  fmt.Printf("The sum of %d + %d is %d \n", x, y, sum)

  swappedX, swappedY := Swap(x, y)
  fmt.Printf("The swap of %d + %d is %d %d \n", x, y, swappedX, swappedY)

  start := 1
  stop := 10
  fmt.Printf("The sum of 1 to 10 with a for loop is %d \n", loops.ForLoop(start, stop))
  fmt.Printf("The sum of 1 to 10 with a while loop is %d \n", loops.WhileLoop(start, stop))

  SqRt(9.0)
}
