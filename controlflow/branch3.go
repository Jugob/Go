package main

import "fmt"

func main(){
	if x:=42; (x>50){
		fmt.Println("x > 50")
	}else if (x<40){
		fmt.Println("x < 40")
	}else{
		fmt.Println("40 < x <50")
	}
}

