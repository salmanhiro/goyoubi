// In this code, we will introduce the math module.

package main

import (
	"fmt"
	//Add the math module
	"math"
)

func main(){
	//The example is pi 
	fmt.Print("Pi approximation: ",math.Pi) //printing number variable, not converted to string. No line break
	fmt.Printf("Pi approximation: %g\n",math.Pi) //printing number variable, converted to string.
	fmt.Println("Pi approximation: ",math.Pi) //printing number variable, not converted to string.
}