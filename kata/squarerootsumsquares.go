package kata

import (
	"math"
	"sort"
)

func ListSquared(m, n int) [][]int {
	l := [][]int{}
	s := sumSliceInt(squares(find_divisors(m, n)))

	keys := make([]int, 0, len(s))
	for k, _ := range s {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := 0; i < len(keys); i++ {
		k := keys[i]
		sqrt := math.Sqrt(float64(s[k]))

		if sqrt == math.Trunc(sqrt) {
			l = append(l, []int{k, s[k]})
		}
	}

	return l
}

func find_divisors(m, n int) map[int][]int {
	p := make(map[int][]int)

	for v := m; v <= n; v++ {
		for w := 1; w <= v; w++ {
			if v%w == 0 {
				p[v] = append(p[v], w)
			}
		}

	}

	return p
}

func squares(m map[int][]int) map[int][]int {
	for k, v := range m {
		for i := 0; i < len(v); i++ {
			m[k][i] = v[i] * v[i]
		}
	}

	return m
}

func sumSliceInt(m map[int][]int) map[int]int {
	n := make(map[int]int)

	for k, v := range m {
		s := 0
		for i := 0; i < len(v); i++ {
			s += v[i]
		}
		n[k] = s
	}

	return n
}
