// This is comment

/*
This 
is 
comment 
too
*/

// The first part of Go code must be package name, and always use main.
package main 

// This part is for importing packages.
// In this case we will import input output called "fmt"
import "fmt"

//The main function of your code
func main(){
	// print "Hello world" in Katakana characters
	fmt.Printf("ハロー・ワールド!")
}

// In your terminal, type `go build 1. Hello World.go` to build the code. 
// To run, just type  1. Hello World.go. 