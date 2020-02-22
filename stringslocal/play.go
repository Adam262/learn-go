package stringslocal

import (
  "fmt"
  "math"
  "strings"

  // Install remote package
  "github.com/google/go-cmp/cmp"
)

func Main() {
  const first, last, alibi = "Ghostface", "Killah", "Ghostdini"
  const chambers = 36

  fmt.Println(ReverseRunes("\noY erehT olleH"))

  fmt.Println(cmp.Diff("Hello World", "Hello Go"))

  fmt.Printf("%s %s is also known as %s.\n", first, last, alibi)
  fmt.Printf("Up from the the %d chambers\n", chambers)

  fmt.Printf("The square root of %d is %g \n", 9, math.Sqrt(9))

  fmt.Printf(
    "The type of `const chambers` %d is %T and the type of `const alibi` %s is %T \n",
    chambers,
    chambers,
    alibi,
    alibi,
  )

  // string replace
  replacer := "foo foo foo"
  fmt.Printf("Replacing foo with bar 2 of 3 times.\n Original: %s \n New: %s \n", replacer, strings.Replace(replacer, "foo", "baz", 2))

  // string join
  sliceToJoin := []string{"a", "b", "c"}
  fmt.Printf("Joining slice ...\n Original: %s \n New: %s \n", sliceToJoin, strings.Join(sliceToJoin, " | "))

}
