package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Moriarity",
		contactInfo: contactInfo{
			zipCode: 945773,
			email:   "jim@email.com",
		},
	}

	fmt.Printf("%+v", jim)
}
