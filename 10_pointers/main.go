package main

import f "fmt"

func main() {
	a := 5
	b := &a // assigning memory address of variable a to b using &
	f.Println(a, b)
	f.Printf("%T\n", a) // checking the type of variable a
	f.Printf("%T\n", b) // checking type :: note the output is *int. Here * means ::::: "pointer"

	// using * to read value from address
	f.Println(*b)

	// change value with the pointer
	*b = 10
	f.Println(a)
}

// Pointers are used to point to memory address of variable
// Other Note :: Every thing in go is passed by value
