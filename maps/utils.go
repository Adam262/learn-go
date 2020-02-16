package maps

import "strings"

func WordCount(s string) (m map[string]int) {
  words := strings.Fields(s)
  m = make(map[string]int)

  for _, v := range words {
    count, ok := m[v]

    if !ok {
      m[v] = 1
    } else {
      m[v] = count + 1
    }
  }

  return
}
