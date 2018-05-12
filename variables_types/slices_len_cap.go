package main

import "fmt"

func main(){
	var countries =[6]string{"Germany","Brazil","England", "Holland","Italy","Argentina"}
	fmt.Printf("Football playing countries:\n%v\n",countries)
	//var  names=make([]string,5,9){"India","Russia","Nepal","Japan"}
	var  names=make([]string,5,9)
	fmt.Printf("names:%v, len=%d, cap=%d\n",names,len(names),cap(names))
	x:=countries[2:4]
	fmt.Printf("name=:%v, len=%d, cap=%d\n",x,len(x),cap(x))
	y:=countries[0:4]
	fmt.Printf("name=:%v, len=%d, cap=%d\n",y,len(y),cap(y))

}
