package kata

import (
	"fmt"
	"strings"
)

func EncryptThis(text string) string {
	if text == "" {
		return ""
	}
	sl_a := strings.Split(text, " ")
	for i := 0; i < len(sl_a); i++ {
		r := []rune(sl_a[i])

		tmp := strings.Split(string(sl_a[i]), "")

		if len(tmp) > 1 {
			t1 := tmp[1]
			t2 := string(tmp[len(tmp)-1])

			tmp[1] = t2
			tmp[len(tmp)-1] = t1
		}

		tmp[0] = fmt.Sprint(r[0])

		sl_a[i] = strings.Join(tmp, "")
	}

	return strings.Join(sl_a, " ")
}
