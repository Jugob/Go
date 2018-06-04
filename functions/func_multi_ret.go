package main

import "fmt"

func calcArea(a float64, b float64) (float64,float64){
	var area, perimeter float64
	area=a*b
	perimeter=2*(a+b)
	return area,perimeter
}

func main(){
	var l,b,area,perimeter float64
	fmt.Println("Enter the length")
        fmt.Scanln(&l)
        fmt.Println("Enter the breadth")
        fmt.Scanln(&b)
	area,perimeter=calcArea(l,b)
	fmt.Println("The area of the rectangle is: ",area)
	fmt.Println("The perimeter of the rectangle is: ",perimeter)
}

