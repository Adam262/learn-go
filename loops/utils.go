package loops

func SumForLoop(start int, stop int) (sum int) {
  sum = 0

  for x := start; x <= stop; x++ {
    sum += x
  }

  return
}

func SumWhileLoop(start int, stop int) (sum int) {
  sum = 0
  x := start

  for x <= stop {
    sum += x
    x += 1
  }

  return
}

func SumRange(coolRange []int) (sum int) {
  sum = 0

  // i is optional index
  for _, v := range coolRange {
    sum += v
  }

  return
}
