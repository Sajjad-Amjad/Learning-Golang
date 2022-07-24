package main

import f "fmt"

func main() {
	// If else
	x := 5
	y := 10

	if x < y {
		f.Printf("%d is less than %d\n", x, y)
	} else {
		f.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "red"

	if color == "red" {
		f.Println("Color is red.")
	} else if color == "blue" {
		f.Println("Color is blue.")
	} else {
		f.Println("Color is NOT red or blue.")
	}

	// Switch
	switch color {
	case "red":
		f.Println("Color is red.")
	case "blue":
		f.Println("Color is blue.")
	default:
		f.Println("Color is NOT blue or red.")

	}

}
