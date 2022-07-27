package main

import (
	f "fmt"
	"strconv"
)

// defone person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " I am " + strconv.Itoa(p.age) + "."
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// Initialize person using struct
	person1 := Person{
		firstName: "Sajjad",
		lastName:  "Amjad",
		city:      "Bijar",
		gender:    "Male",
		age:       16,
	}
	f.Println(person1)
	// Alternative :: Asign Values without using properties name
	person2 := Person{"Sajjad", "Amjad", "Bijar", "Male", 16}
	f.Println(person2)
	f.Println(person2.firstName) // Get single field :: firstName
	person2.age++
	f.Println(person2)
	f.Println("----------------------------")

	person1.hasBirthday()
	f.Println(person1.greet())

}

//  Structs are like classes in golang.
// NOTE:: there is no classes in golang

////////////
