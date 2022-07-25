package main

import f "fmt"

func main() {
	ids := []int{22, 4, 5, 5, 6, 4, 3}

	// Loops through ids
	for i, id := range ids {
		f.Printf("%d - ID: %d\n", i, id)
	}

	// Not usig index :: i
	for _, id := range ids {
		f.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	f.Println("Sum", sum)

	// range with maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com", "Mike": "mike@gmail.com"}
	for k, v := range emails {
		f.Printf("%s: %s\n", k, v)
	}

	// printing onlyy keys :: k :: names
	for k := range emails {
		f.Println("Name: " + k)
	}
}

// Range is used to loop over arrays, slices, maps, etc
