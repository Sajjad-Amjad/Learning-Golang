package main

import f "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// // Assign Key : Values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// declare and add kv at same time
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com", "Mike": "mike@gmail.com"}

	f.Println(emails)      // printing map
	f.Println(len(emails)) // getting length of map
	f.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	f.Println(emails)
}
