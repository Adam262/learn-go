package main

import (
	"fmt"
	"math"

	// Install local package

	"github.com/user/hello/loops"
	"github.com/user/hello/mathstuff"
	"github.com/user/hello/morestrings"

	// Install remote package
	"github.com/google/go-cmp/cmp"
)

func main() {
	const first, last, alibi = "Ghostface", "Killah", "Ghostdini"
	const chambers = 36

	fmt.Println(morestrings.ReverseRunes("\noY erehT olleH"))

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	fmt.Printf("%s %s is also known as %s.\n", first, last, alibi)
	fmt.Printf("Up from the the %d chambers\n", chambers)

	fmt.Printf("The square root of %d is %g \n", 9, math.Sqrt(9))

	// var keyword variable declaration with explicit type + initial value
	var x int = 2

	// var keyword variable declaration with explicit type but no initial value
	var y int
	y = 5

	// short form variable declaration with implicit type.
	// short form is only available within a function
	sum := mathstuff.Add(x, y)
	fmt.Printf("The sum of %d + %d is %d \n", x, y, sum)

	swappedX, swappedY := mathstuff.Swap(x, y)
	fmt.Printf("The swap of %d + %d is %d %d \n", x, y, swappedX, swappedY)

	start := 1
	stop := 10
	fmt.Printf("The sum of 1 to 10 with a for loop is %d \n", loops.ForLoop(start, stop))
	fmt.Printf("The sum of 1 to 10 with a while loop is %d \n", loops.WhileLoop(start, stop))

	mathstuff.SqRt(9.0)
}
