package pointers

import "fmt"

// Go has pointers. A pointer holds the memory address of a value.
// The type *T is a pointer to a T value. Its zero value is nil.
func Main() {
	var p *int

	// The & operator generates a pointer to another variable.
	i := 42
	p = &i

	// The * operator denotes the pointer's underlying value.
	// This is known as "dereferencing" or "indirecting".
	fmt.Println("*p is a pointer to value i. setting i to 42. now *p is dereferenced to: ", *p) // read i through the pointer p

	*p = 21 // set i through the pointer p
	fmt.Println("setting *p to 21. this will set the value of i. now i is: ", i)
}
