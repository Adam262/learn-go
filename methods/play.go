package methods

import "fmt"

func Main() {
  var floaty Floaty
  floaty = -1.23
  fmt.Printf("Running method on Floaty receiver. The absolute value of floating number %v is %v \n", floaty, floaty.Abs())

  usain := &Person{FirstName: "Usain", LastName: "Bolt"}
  fmt.Printf("Running method on *Person receiver. The person's name is %s \n", usain.FullName())
}

// struct receiver
type Person struct {
  FirstName string
  LastName  string
}

func (p *Person) FullName() string {
  return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// non struct receiver
type Floaty float64

func (f Floaty) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }

  return float64(f)
}
