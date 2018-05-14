package main

import "fmt"

func main(){
	var months map[int]string
	//months=make(map[int]string)
	months[1]="January"
	months[2]="February"
	months[3]="March"

	for mon:=range months{
		fmt.Println("Month number ",mon, "is ",months[mon])
	}
}



