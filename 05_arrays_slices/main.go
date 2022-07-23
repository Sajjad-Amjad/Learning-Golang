package main

import f "fmt"

func main() {
	//Arrays
	var fruitArr [2]string
	//Assign Values
	fruitArr[0] = "Apples"
	fruitArr[1] = "Oranges"
	f.Println(fruitArr)
	f.Println(fruitArr[0])
	f.Println(fruitArr[1])

	f.Println("---------------------------------------")

	// Declaring and assigning values to array at same
	fruitArr1 := [2]string{"Mango", "Banana"}
	f.Println(fruitArr1)

	f.Println("---------------------------------------")
	// Slices
	fruitSlice := []string{"Mango", "Banana", "Cerry", "Grape"}
	f.Println(fruitSlice)
	f.Println("Length of Slice: ", len(fruitSlice))
	f.Println("Index 1 to 3: ", fruitSlice[1:3])
}
