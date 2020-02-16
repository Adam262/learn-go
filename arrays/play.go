package arrays

import "fmt"

func Main() {
  // arrays are fixed size
  var a [2]string
  a[0] = "zero"
  a[1] = "one"

  primes := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

  fmt.Println(a)
  fmt.Println(primes)

  // A slice is a dynamically-sized, flexible view into the elements of an array.
  // Slices are much more common than arrays.
  // The type []T is a slice with elements of type T.
  // Slices do not store data - they are references to an array

  slice := primes[0:3]
  fmt.Println(slice)

  // can build slice literal without size
  newSlice := []bool{true}
  // newSlice[1] = false
  fmt.Println(newSlice)
}
