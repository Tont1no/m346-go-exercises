package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}

	// TODO: declare a type for Class (consisting of multiple students)
	type Class []Student

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[string]Class{
		"Math": {
			{FirstName: "Alice", LastName: "Smith"},
			{FirstName: "Bob", LastName: "Johnson"},
		},
		"Science": {
			{FirstName: "Charlie", LastName: "Brown"},
			{FirstName: "Dana", LastName: "White"},
		},
		"History": {
			{FirstName: "Eve", LastName: "Black"},
		},
	}

	// TODO: output everything using fmt.Println()
	for module, class := range modules {

		fmt.Println("Module:", module)
		for _, student := range class {
			fmt.Printf("- %s %s\n", student.FirstName, student.LastName)
		}

	}

}
