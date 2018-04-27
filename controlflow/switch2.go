package main

import (
	"fmt"
	"time"
)
func main(){
	fmt.Println(time.Now().Weekday())
	switch (time.Now().Weekday()){
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thursday")
	case time.Friday:
		fmt.Println("Friday")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	}
}

