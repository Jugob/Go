package main

import ( 
	"fmt"
)
func main(){
	var firstvar,secvar int
	fmt.Println("Please enter the first value")
	fmt.Scanln(&firstvar)
	fmt.Println(firstvar)
	fmt.Println("Please enter the second value")
        fmt.Scanln(&secvar)
	fmt.Println(secvar)
}

