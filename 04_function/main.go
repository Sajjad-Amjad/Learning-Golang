package main

import f "fmt" // f"fmt"  means that import fmt as f :: same as python

func greetings(name string) string {
	return "Hello " + name // string concentation
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	f.Println(greetings("Sajjad"))
	f.Println(getSum(3, 4))
}
