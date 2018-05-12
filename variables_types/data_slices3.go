package main

import "fmt"

func main(){
	var ar [10]int
	fmt.Println("ar=",ar)
	sl:=ar[:6]
	fmt.Println("sl=",sl)
	ar[4]=100
	fmt.Println("ar=",ar)
	fmt.Println("sl=",sl)
}
