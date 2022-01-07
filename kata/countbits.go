package kata

import (
	"strconv"
	"strings"
)

func countBits(x uint) int {
	i := uint64(x)
	b := strconv.FormatUint(i, 2)
	return strings.Count(b, "1")
}
