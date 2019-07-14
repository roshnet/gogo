package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// sherlock := person{firstName: "Sherlock", lastName: "Holmes"}

	var sherlock person // Allows zero values to assign
	sherlock.firstName = "Sherlock"
	sherlock.lastName = "Holmes"
	fmt.Println(sherlock)
	fmt.Printf("%+v", sherlock)
}
