package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Introduction to Programming",
		117: "Data Structures and Algorithms",
		346: "Advanced Go Programming",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 117)
	// TODO: add one
	modules[212] = "Database Systems"
	// TODO: replace one
	modules[346] = "Concurrent Programming in Go"

	fmt.Println(modules)
}
