package mathlocal

import (
  "fmt"
  "math"
)

// can also do func Add(x, y int) bc they are same data type
// note named return value
func Add(x int, y int) (sum int) {
  sum = x + y
  return
}

// func can return two results
func Swap(x, y int) (int, int) {
  return y, x
}

func SqRt(x float64) (squareRoot float64) {
  z := 1.0

  for i := 0; i < 10; i++ {
    z -= (z*z - x) / (2 * z)

    fmt.Println(z)
  }

  return
}

func Hyp(a, b float64) (c float64) {
  c = math.Sqrt(a*a + b*b)
  return
}
