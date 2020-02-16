package maps

import (
  "fmt"

  "golang.org/x/tour/wc"
)

type Vertex struct {
  Lat, Long float64
}

func Main() {
  // build map via make
  vm := make(map[string]Vertex)

  vm["somePlace"] = Vertex{1.23, 4.56}

  fmt.Printf("This map %v is made via `make(map[string]Vertex)`\n", vm)

  // build map via map literal
  mapLiteral := map[string]Vertex{
    "somePlace":      {1.23, 4.56},
    "someOtherPlace": {5.78, 1.23},
  }

  // add key
  mapLiteral["yoAnotherPlace"] = Vertex{1.25, 3.57}

  fmt.Printf("This map %v is made via `mapLiteral := map[string]Vertex{k:v, k:v,}` \n", mapLiteral)

  // get value
  value, ok := mapLiteral["yoAnotherPlace"]
  fmt.Printf("The map has value yoAnotherPlace %g and ok returns %t \n", value, ok)

  value, ok = mapLiteral["notAPlace"]
  fmt.Printf("The map has no value notAPlace %g and ok returns %t \n", value, ok)

  // test WordCount
  fmt.Println("Testing WordCount util...")
  wc.Test(WordCount)
}
