package main

import (
	"fmt"
	"math" //importing multiple packages in golang
	"mymodule"
)

// import "fmt"  // to import a single package at a time
func main() {
	fmt.Println(math.Floor(2.67))
	fmt.Println(math.Ceil(2.67))
	fmt.Println(mymodule.Reverse("olleh"))
}
