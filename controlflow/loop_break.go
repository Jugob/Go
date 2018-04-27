package main

import "fmt"

func main(){
	for i:=1;i<100;i=i+2{
		if i%5==0 && i%3==0{
			break
		}
		fmt.Println(i)
	}
}
