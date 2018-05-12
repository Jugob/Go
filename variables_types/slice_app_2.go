package main

import "fmt"

func main(){
	var ar1 [6]int
	fmt.Printf("ar1=:%v, len=%d, cap=%d\n",ar1,len(ar1),cap(ar1))
	sl1:=ar1[:4]
	fmt.Printf("sl1=:%v, len=%d, cap=%d\n",sl1,len(sl1),cap(sl1))
	sl1=append(sl1, 1,2,3,4,5,6,7,8,9)
	fmt.Printf("sl1=:%v, len=%d, cap=%d\n",sl1,len(sl1),cap(sl1))
	fmt.Printf("ar1=:%v, len=%d, cap=%d\n",ar1,len(ar1),cap(ar1))
	//When the slice has exceeded the lgnth of the underlying array
	//slice is reallocated.
	sl1[0]=444
	fmt.Printf("sl1=:%v, len=%d, cap=%d\n",sl1,len(sl1),cap(sl1))
	fmt.Printf("ar1=:%v, len=%d, cap=%d\n",ar1,len(ar1),cap(ar1))
}
