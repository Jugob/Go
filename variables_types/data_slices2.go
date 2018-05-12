package main

import "fmt"

//This program shows that make() is not really effective with slices.
//make() does not do anything that cannot be done without it.

func main(){
	ar:=[10]int{1,2,3,4,5,6,7,8}    //array
	fmt.Println("ar=",ar)		//Print array
	fmt.Println("ar[:5]=",ar[:5])	//slice an array and Print
	sl:=make([]int, 10)		//create a slice(array)using make()
	fmt.Println("sl=",sl)		//Print slice(array)
	sl[1]=100
	fmt.Println("sl=",sl)
	fmt.Println("sl[:5]=",sl[:5])	//Print slice of a slice(array)

}

