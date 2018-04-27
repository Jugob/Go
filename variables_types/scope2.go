package main

import "fmt"

var x string = "Hello World"
var y int
func main(){

fmt.Print("Hello\n")
/* see the difference between two printing styles. fmt.Println is just like fmt.Print("\n") */
fmt.Println(x)

/*There is nothing like an uninitialized variable in Go. "int" variable type is initialized with value 0, "string" variable is left blank. It is also called "zero value variable" */
fmt.Println(y)

/*we can change the values assigned to a variable. */
x= "bye bye world"
fmt.Println(x)

/* although we have declared the variable type in the beginning of the program, it is not necessary to explicitly declare the type of the variable. the variable is also implicitly decided by the value assigned to the variable. Example : */

z:="new variable"
/* This is called the shorthand for  variable creation and assignment.
This is equivalent to 
var z string = "new variable"
The catch is: This can only be done inside a function, not outside.
However, the following cannot be done:

z:=7
z:=8

this will throw and error.
If the value of the value has to be changed, it must be regular assignment
z=8
shorthand can only be used once for a variable.

*/

fmt.Println(z)

/* However, Go is static type and strong typed. Once a variable acquires a value and hence a type, or just a type, it cannnot be assigned a value of a different type. Example: 

z:=5 
OR
z=5
fmt.Println(z)

this would throw an error.

*/

/* Using a variable before it is defined, will throw and error. However, if a variable is defined outside of a function, */


/* In Go, you cannot  create a variable and NOT not use it. This will throw and error. It is quite obviously a mistake, otherwise why would one create a variable and not use it, it is a waste of resource*/

/*

Scope:
Easy to remember: scope of a variable ends with the ending curly braces inside which the variable has been created. If the curly braces are nested, then the scope of the variable declared outside the nested brace is still valid in the nested brace.


*/

/* Semi colon:
In one word: It is not needed.
However, when you use more than one statement in a line, you can use semi colon to separate the statements.

*/





}
