package main

import (
  "fmt"
  "os"

/* Go forces you to use an imported package. If you import and not use a package, then it is an error in the go program.*/
)

const (
  x=10
  y=10
)

/* usually constants are declared at the beginning, outside of any function. */
func main(){

 const x=5
 fmt.Println(x)

 fmt.Println(y)
 

}
