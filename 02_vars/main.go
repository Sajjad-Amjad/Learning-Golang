package main

import "fmt"

func main(){
	// MAIN TYPES
    // string
	// bool
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint31 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Creating Variables
	var name string = "Sajjad"   // var name = "Sajjad" :: same thing  :: go will auto identify the type
	var age int = 16             // var age = "16" :: same thing :: go will auto identify the type
	const name2 string = "Sajjad" 
	var age2 int32 = 16
	var isCOOL bool = true
	var size  float32 = 1.3
	fmt.Println(name, age, name2, age2, isCOOL, size)


	// Shorthand for creating variables
	name3 := "Sajjad Amjad"
	size2 := 1.4  // default the type is float64 
	fmt.Println("NAME:", name3)


	fmt.Printf("%T\n", name)  // To print the type of variable
	fmt.Printf("%T\n", age)  // To print the type of variable
	fmt.Printf("%T\n", name2)  // To print the type of variable
	fmt.Printf("%T\n", isCOOL)  // To print the type of variable
}

// NOTE:
//      name := "sajjad"  ::: this type of variable declaration will not work outside the function