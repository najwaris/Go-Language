// const declares a constant value.
// A const statement can appear anywhere a var statement can.
// Constant expressions perform arithmetic with arbitrary precision.
// A numeric constant has no type until it’s given one, such as by an explicit conversion.
// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. 
// For example, here math.Sin expects a float64.

// Code starts below:

package main

// Import fmt and math
import (
    "fmt"
    "math"
)

// This is a constant string declaration
const s string = "constant" 

func main() {
  
    // Print the value of the constant string
    fmt.Println(s) 

    // Declare a constant integer
    const n = 500000000 

    // Calculate the value of a constant floating-point number
    const d = 3e20 / n 

    // Print the value of the constant floating-point number
    fmt.Println(d) 

    // Convert the constant floating-point number to an int64 and print it
    fmt.Println(int64(d)) 

    // Calculate and print the sine of the constant integer
    fmt.Println(math.Sin(n)) 
}
