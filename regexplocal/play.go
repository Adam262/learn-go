package regexplocal

import (
	"fmt"
	"regexp"
)

func Main() {
	example := "foo,    baz   , this is bar"

	fmt.Printf("Cleaning example:  `%s`. Resulting regexp is: %v \n", example, cleaner(example))
}

func cleaner(commaSeparatedStr string) *regexp.Regexp {
	cleanerRegExp := regexp.MustCompile(`(\s*\,\s*)`)

	return regexp.MustCompile(
		fmt.Sprintf(
			"^(%s)$",
			cleanerRegExp.ReplaceAllString(commaSeparatedStr, "|"),
		),
	)
}
