package main

import "fmt"

type person struct{
	firstName string
	lastName string
	contact contactInfo
}

type contactInfo struct{
	email string
	zipcode int
}

func main(){

	alex := person{firstName: "Alex", 
	lastName: "Anderson",
	contact: contactInfo{
		email: "alex@gmail.com",
		zipcode: 78984,
	}}
	fmt.Println(alex)

	var ram person
	ram.firstName = "Ram"
	ram.lastName = "Das"
	ram.contact.email ="ram@gmail.com"
	ram.contact.zipcode = 79964

	fmt.Println(ram)

}