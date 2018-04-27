package main

import "fmt"

func main(){
	var sum int=0
	for i:=0;i<10;i++{
		sum=sum+i
	}
	fmt.Println("The sum is: ",sum)
        //use as a while loop
	j:=1
	for j<=5{
		fmt.Println(j)
		j=j+1
	}

	//for loop alos can be used as an infinite loop
	for{
		//do something
		//This will keep running until keyboaed interrupt.
	}
}

