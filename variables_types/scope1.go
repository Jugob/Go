package main

import "fmt"

var x string
var y int
func main(){

fmt.Print("Hello\n")
fmt.Print(x)

/*There is nothing like an uninitialized variable in Go. "int" variable type is initialized with value 0, "string" variable is left blank.*/
fmt.Print(y)
}
