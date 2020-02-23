package interfaces

import (
  "errors"
  "fmt"
  "math"
  "reflect"
)

func Main() {
  coolCircle := &Circle{radius: 3.0}

  radRect := &Rect{length: 3.0, width: 4.0}

  measure(coolCircle)
  measure(radRect)

  fmt.Println("An unassigned interface will trigger a nil pointer error unless `measure()` handles this case")
  var g geometry
  measure(g)

  dog := Dog{}
  dog.Age = 3
  fmt.Printf("Dog.Age is defined as an empty interface.\n The dog has age %v and type %T \n", dog.Age, dog.Age)

  dog.Age = "3"
  fmt.Printf("The dog has age %v and type %T \n", dog.Age, dog.Age)

}

type geometry interface {
  perim() float64
  area() float64
}

func measure(g geometry) (err error) {
  fmt.Printf("This geometry has type %T has value %v \n", g, g)

  if isNil(g) {
    err = errors.New("Error: geometry cannot be nil")
    fmt.Println(err)
    return
  }

  fmt.Printf("This geometry has perim %f and area %f \n", g.perim(), g.area())

  return
}

func isNil(i interface{}) bool {
  return i == nil || reflect.ValueOf(i).IsNil()
}

type Circle struct {
  radius float64
}

func (c *Circle) perim() float64 {
  return 2 * math.Pi * c.radius
}

func (c *Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

type Rect struct {
  length float64
  width  float64
}

func (r *Rect) perim() float64 {
  return (2 * r.length) + (2 * r.width)
}

func (r *Rect) area() float64 {
  return r.length * r.width
}

// empty interface
type Dog struct {
  Age interface{}
}
