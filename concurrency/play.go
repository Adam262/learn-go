package concurrency

import (
  "fmt"
)

func Main() {
  slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
  ch := make(chan int)

  go concurrentSum(slice[:(len(slice)/2)], ch)
  go concurrentSum(slice[(len(slice)/2):], ch)

  x, y := <-ch, <-ch

  fmt.Printf(
    "Summing a slice of ints concurrently using 2 channels...\n Input: %v \n Ch1 Output: %v Ch2 Output: %v Sum: %v \n ",
    slice,
    x,
    y,
    x+y,
  )
}

func concurrentSum(inputSlice []int, c chan int) {
  sum := 0

  for _, v := range inputSlice {
    sum += v
  }

  c <- sum
}
