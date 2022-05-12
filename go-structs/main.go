package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

// nested struct
type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func main() {
	fmt.Println("Structs!")

	//engineer := Engineer{Name: "Paul"}

	//verbose approach
	engineer := Engineer{
		Name: "Paul",
		Age:  41,
		Project: Project{
			Name:         "Go structs!",
			Priority:     "High",
			Technologies: []string{"Go"},
		},
	}

	//fmt.Println(engineer)

	// to see values and field names
	fmt.Printf("%+v\n", engineer)
}
