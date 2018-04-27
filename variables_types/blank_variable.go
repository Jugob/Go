package main

import (
   "fmt"
   _ "os"
)

/* When a package is imported and not used, it means an error in Go. to avoid the error, we can use a blank identifier. The same is the case with the variables also.

There are situations when you would like to use a blank variable.
a,_ = func1()
if the func1 returns more than 1 value, then you store it in the blank identifier.

 */



func main(){
 fmt.Println("Hello")




}
