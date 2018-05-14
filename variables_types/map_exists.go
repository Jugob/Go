package main

import "fmt"

func main(){
	var months map[int]string
        //The following line actually creates the map object
        months=make(map[int]string)
        months[1]="January"
        months[2]="February"
        months[3]="March"
        for mon:=range months{
		fmt.Println("Month number ",mon, "is ",months[mon])
	}
	val,exist:=months[2]
	if exist{
		fmt.Println(val," exists")
	}else{
		fmt.Println("Does not exist")
	}

	val,exist=months[5]
	if exist{
		fmt.Println(val," exists")
	}else{
		fmt.Println("Does not exist")
	}

}
