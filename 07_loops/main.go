package main

import f "fmt"

func main() {
	i := 1
	// Long Method
	for i <= 10 {
		f.Println(i)
		i++ // or i = i + 1
	}

	// Short Method
	for i := 1; i <= 10; i++ {
		f.Printf("Number %d\n", i)
	}

	// FizzBuzz ::: Challenge
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			f.Println("FizzBuzz")
		} else if i%3 == 0 {
			f.Println("Fizz")
		} else if i%5 == 0 {
			f.Println("Buzz")
		} else {
			f.Println(i)
		}
	}
}
