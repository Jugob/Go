package main

import "fmt"

func main(){
	var ar1= [10]int{1,2,3,4,5,6,7,8}
	var ar2 [10]int
	sl1 :=ar1[:10]
	sl2 :=ar1[:10]
	ar2=ar1
	fmt.Println("ar1=",ar1, "ar2=", ar2, "sl1=",sl1, "sl2=", sl2)
	//fmt.Println("&ar1=",&ar1, "&ar2=", &ar2, "&sl1=",&sl1, "&sl2=", &sl2,&ar1, &ar2,&sl1, &sl2)
	fmt.Printf("&ar1=%p, &ar2=%p, &sl1=%p, &sl2=%p \n", &ar1, &ar2,&sl1, &sl2)
}


