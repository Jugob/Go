package main

import "fmt"

func main(){
	var countries =[6]string{"Germany","Brazil","England", "Holland","Italy","Argentina"}
	fmt.Printf("Football playing countries:\n%v\n",countries)
	x:=countries[:4]
	fmt.Printf("x=:%v, len=%d, cap=%d\n",x,len(x),cap(x))
	y:=countries[2:6]
	fmt.Printf("y=:%v, len=%d, cap=%d\n",y,len(y),cap(y))
	copy(y,x)
	fmt.Printf("x=:%v, len=%d, cap=%d\n",x,len(x),cap(x))
	fmt.Printf("y=:%v, len=%d, cap=%d\n",y,len(y),cap(y))
	fmt.Printf("countries=:%v\n",countries)

}

