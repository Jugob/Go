package main

import "fmt"

func main(){
	var myIntArray [5]int
	var myStrArray [5]string
	myIntArray[3]=42
	myStrArray[3]="Hello"
	myStrArray[4]="World"
	fmt.Println(myIntArray)
	fmt.Println(myStrArray)
}
