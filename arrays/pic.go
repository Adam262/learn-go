package arrays

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (outerSlice [][]uint8) {
  outerSlice = make([][]uint8, dy, dy)
  innerSlice := make([]uint8, dx, dx)

  for i, _ := range outerSlice {
    for j, _ := range innerSlice {
      innerSlice = append(innerSlice, transform(i, j))
    }
  }

  return
}

func transform(x, y int) uint8 {
  return uint8(x * y)
}

func main() {
  pic.Show(Pic)
}
