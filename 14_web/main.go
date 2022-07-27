package main

import (
	f "fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	f.Fprintf(w, "<br><h1>Hello WOrld</h1>") //# to write in browser
	// f.Fprintf(w, "<br><h1>Hello WOrld</h1>") //# to write in browser

}

func main() {
	http.HandleFunc("/", index)
	f.Println("Server Started.....")
	http.ListenAndServe(":3000", nil)
}
