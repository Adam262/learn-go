package main

import (
	"fmt"

	// Install local package
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
	fmt.Printf("Up from the the %d chambers", chambers)
}