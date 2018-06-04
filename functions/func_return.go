package main

import "fmt"

func calcArea(a float64, b float64) float64{
	return a*b
}

func main(){
	var l,b,area float64
	fmt.Println("Enter the length")
        fmt.Scanln(&l)
        fmt.Println("Enter the breadth")
        fmt.Scanln(&b)
	area=calcArea(l,b)
	fmt.Println("The area of the rectangle is: ",area)
}

