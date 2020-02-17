package funcs

import (
  "fmt"
)

func Main() {
  values := []int{1, 2, 3, 4, 5}

  fmt.Printf(
    "Mapping slice `values` %v to a slice of squares %v \n",
    values,
    MapInts(values, square),
  )

  fmt.Println("Running adder loop...", adderLoop(10))

  fmt.Println("Calculating Fibonacci...", Fib(10))
}

// Higher order function takes function as argument
func MapInts(input []int, fn func(x int) (y int)) (output []int) {
  output = make([]int, len(input), cap(input))

  for i, v := range input {
    output[i] = fn(v)
  }

  return
}

func square(x int) int { return x * x }

func Adder() func(int) int {
  sum := 0

  return func(x int) int {
    sum += x
    return sum
  }
}

// example of closure
func adderLoop(length int) (sum int) {
  newAdder := Adder()
  sum = 0

  for x := 1; x <= length; x++ {
    sum = newAdder(x)
  }

  return
}

func Fib(length int) (fibs []int) {
  x := 0
  y := 1
  fibs = []int{x, y}

  for current := x; current <= length; current++ {
    sum := x + y

    fibs = append(fibs, sum)

    x = y
    y = sum
  }

  return
}
