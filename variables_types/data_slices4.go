package main

import "fmt"

func main(){
	x:=make([]float64,6)
	fmt.Println(x)
	slice2:=append(x, 2,3,4)
	fmt.Println(slice2)
}
