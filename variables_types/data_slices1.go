package main

import (
	"fmt"
	"reflect"
)

func main(){
	var x= [5]float64{5,4,6,2.6,7.9}
	fmt.Println(x)
	var y=x[1:4]
	fmt.Println(y)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))
}
