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

	fmt.Printf("%+v", ram)

	ram.updateName("Gopal")
	fmt.Println(ram) // no change in main object

	ramPointer := &ram
	ramPointer.updateFirstName("Gopal")
	fmt.Println(ram)

	ram.updateFirstName("Kesahv")
	fmt.Println(ram)

}

func (p person) updateName(newFirstName string){
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateFirstName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}
