package loops

import "fmt"

func Main() {
  start := 1
  stop := 10
  fmt.Printf("The sum of 1 to 10 with a for loop is %d \n", SumForLoop(start, stop))
  fmt.Printf("The sum of 1 to 10 with a while loop is %d \n", SumWhileLoop(start, stop))

  values := []int{1, 2, 3}
  fmt.Printf("The sum of values %v with a range is %d \n", values, SumRange(values))
}
