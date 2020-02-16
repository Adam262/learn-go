package structs

import "fmt"

type Athlete struct {
  FirstName   string
  LastName    string
  Nationality string
}

func Main() {
  a := Athlete{"Usain", "Bolt", ""}
  p := &a
  p.Nationality = "Jamaican"
  fmt.Printf("The person's name is %s %s and he is %s \n", p.FirstName, p.LastName, p.Nationality)
}
