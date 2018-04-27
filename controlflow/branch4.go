package main

import ( 
	"fmt"
)
//the program below checks if a number is divisible by 3 OR 5 and 7 AND 11.
//Prints an appropriate message.
func main(){
	var usernum int
	fmt.Println("Please enter the first value")//prompt to the user
	fmt.Scanf("%d", &usernum) //read the user input
	if (usernum%3==0) || (usernum%5==0){
		fmt.Printf("The number %d is divisible by 3 or 5\n", usernum)
	}
	if (usernum%11==0) && (usernum%7==0){
		fmt.Printf("The number %d is divisible by 11 and 7\n", usernum)
	}
}
