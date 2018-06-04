package main

import "fmt"
//a function that takes 2 inputs of type int each and returns their sum.
func add(a int,b int) Nil {
	return a+b
}
//main function to read inputs from the user
func main(){
	var a, b, sum int
	fmt.Println("Enter a number")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number")
	fmt.Scanln(&b)
	sum=add(a,b)
	fmt.Println("The sum is:",sum)
}
