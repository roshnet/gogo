package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Moriarity",
		contact: contactInfo{
			zipCode: 945773,
			email:   "jim@email.com",
		},
	}

	fmt.Printf("%+v", jim)
}
