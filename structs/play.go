package structs

import "fmt"

type Athlete struct {
  FirstName   string
  LastName    string
  Nationality string
}

func (a *Athlete) FullName() string {
  return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}

func Main() {
  a := Athlete{"Usain", "Bolt", ""}
  p := &a
  p.Nationality = "Jamaican"
  fmt.Printf("The person's name is %s and he is %s \n", p.FullName(), p.Nationality)
}
