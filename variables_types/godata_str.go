package main

import (
	"fmt"
	"reflect"
)

func main(){
	var hello string ="Hello"
	fmt.Println(reflect.TypeOf(hello))
	fmt.Printf("\nType= %T\n", hello)
}
