package main

import "fmt"

func main(){
	a:=[]string{"a","b","c","d"}
	for i:=range a{
		fmt.Println(i)
	}

	alpha:=map[string] string {"a":"apple", "b":"banana", "c":"cherry"}
	for key := range alpha{
		fmt.Println("The keys - value pair in alpha is:",key,":",alpha[key])
	}
}
