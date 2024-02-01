// Strings, which can be added together with +.
// Integers and floats.
// Booleans, with boolean operators as you’d expect.

//CODE STARTS BELOW: 

// values.go
package main

// import package "fmt"
import "fmt"

// main function
func main() {

	// Concatenating strings
	fmt.Println("go" + "lang") 

	// Performing addition
	fmt.Println("1+1 =", 1+1) 

	// Performing subtraction
	fmt.Println("7.0/3.0 =", 7.0/3.0) 

	// Logical AND operation
	fmt.Println(true && false) 

	// Logical OR operation
	fmt.Println(true || false) 

	// Logical NOT operation
	fmt.Println(!true) 
}
