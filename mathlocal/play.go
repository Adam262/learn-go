package mathlocal

import (
  "fmt"
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

  SqRt(9.0)
}
