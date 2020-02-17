package arrays

import "fmt"

func Pic(dx, dy int) (outerSlice [][]uint8) {
  outerSlice = make([][]uint8, dy, dy)
  innerSlice := make([]uint8, dx, dx)

  for i, _ := range outerSlice {
    for j, _ := range innerSlice {
      fmt.Println("i: ", i)
      fmt.Println("j: ", j)
      innerSlice = append(innerSlice, transform(i, i))
    }
  }

  return
}

func transform(x, y int) uint8 {
  return uint8(x * y)
}
