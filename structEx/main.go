package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	defaultPerson := person{
		firstName: "default",
		lastName:  "banerjee",
		contact: contactInfo{
			email: "default@mail.com",
			zip:   110111,
		},
	}


	personPointer := &defaultPerson
	personPointer.updateName("thanos")
	defaultPerson.printPerson()

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) printPerson() {
	fmt.Printf("%+v", p)
}
