package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func str(x, y string) string {
	return x +" "+y
}
var(
	// Integers
	a int = 12
	b int = 11
	// Strings
	c string = "Bella"
	d string = "Ciao"

)
func main() {
	//Addition of integer
	fmt.Println(add(a, b))
	//Addition of string
	fmt.Println(str(c, d))
}