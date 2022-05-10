package main

import "fmt"

type Engineer struct {
	Name string
	Age  int
}

// method to update struct - pointer reciever

func (e *Engineer) UpdateAge() {
	e.Age += 1
}

// value reciever - copy of address, not original

func (e Engineer) UpdateName() {
	e.Name = "new name"
}

func main() {
	fmt.Println("Pointers with struct!")

	paul := &Engineer{
		Name: "Paul",
		Age:  41,
	}
	fmt.Println(paul)

	// update
	paul.UpdateAge()
	fmt.Println(paul)

	// value reciever uses copy of address
	paul.UpdateName()
	fmt.Println(paul)
}

/*
func main() {
	fmt.Println("Go Pointers!!")

	var name string
	name = "Paul"
	fmt.Println(name)

	ptr := &name
	fmt.Println(*ptr)

	// update value stored
	*ptr = "Conor"
	fmt.Println(name)
}
*/
