package arrays

import (
  "fmt"
)

func Main() {
  // arrays are fixed size
  var a [2]string
  a[0] = "zero"
  a[1] = "one"
  fmt.Printf("array `a` %v has fixed length 2\n", a)

  // can also let compiler infer size
  var b = [...]string{"two", "three"}
  fmt.Printf("array `b` %v also has fixed length 2\n", b)

  primes := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

  // A slice is a dynamically-sized, flexible view into the elements of an array.
  // Slices are much more common than arrays.
  // The type []T is a slice with elements of type T.
  // Slices do not store data - they are references to an array

  slice := primes[0:3]
  fmt.Printf("`slice = primes[0:3]` %v is a reference to array `primes`\n", slice)

  // Can omit high / low bounds, so below are equivalent
  // primes[0:10]
  // primes[:10]
  // primes[0:]
  // primes[:]
  fullSlice := primes[:]
  fmt.Printf("A full slice of array primes[10]int` %v can be written as `primes[:], primes[0:], primes[:10] or primes[0:10]`\n", fullSlice)

  // can build slice literal without explcit size
  newSlice := []bool{true}
  fmt.Printf(
    "`newSlice := []bool{true}` %v is a slice literal of type %T with length: %d, capacity: %d\n",
    newSlice,
    newSlice,
    len(newSlice),
    cap(newSlice),
  )
  // newSlice[1] = false

  makeSlice := make([]bool, 0, 5)
  fmt.Printf(
    "`makeSlice := make([]bool, 0, 5)` %v is a slice literal of type %T with length: %d, capacity: %d\n",
    makeSlice,
    makeSlice,
    len(makeSlice),
    cap(makeSlice),
  )

  // Appending to a slice creates a new slice
  makeSlice = append(makeSlice, true)
  makeSlice = append(makeSlice, false)

  fmt.Printf(
    "`makeSlice := make([]bool, 0, 5)` %v is a slice literal of type %T with length: %d, capacity: %d\n",
    makeSlice,
    makeSlice,
    len(makeSlice),
    cap(makeSlice),
  )

  // pic.Show(Pic)
}
