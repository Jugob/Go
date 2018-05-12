package main

import "fmt"

func main(){
	var ar1 [6]int
        fmt.Printf("ar1=:%v, len=%d, cap=%d\n",ar1,len(ar1),cap(ar1))
        sl1:=ar1[:4]
        fmt.Printf("sl1=:%v, len=%d, cap=%d\n",sl1,len(sl1),cap(sl1))
	sl2:=make([]int,5,2*len(ar1))
	copy(sl1,sl2)
        fmt.Printf("sl1=:%v, len=%d, cap=%d\n",sl1,len(sl1),cap(sl1))
        fmt.Printf("sl1=:%v, len=%d, cap=%d\n",sl2,len(sl2),cap(sl2))
}

