// Variables are explicitly declared and used by the compiler
// You can declare multiple variables at once.
// Go will infer the type of initialized variables.
// Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. 
// This syntax is only available inside functions.

// Code starts below: 

package main

// import package "fmt"
import "fmt"

func main() {

    // Initializing variable 'a' with the value "initial"
    var a = "initial" 
    fmt.Println(a)

    // Initializing variables 'b' and 'c' with the values 1 and 2 respectively
    var b, c int = 1, 2 
    fmt.Println(b, c)

    // Initializing variable 'd' with the value true
    var d = true 
    fmt.Println(d)

    // Declaring variable 'e' without initializing it
    var e int 
    fmt.Println(e)

    // Short variable declaration, initializing variable 'f' with the value "apple"
    f := "apple" 
    fmt.Println(f)
}
