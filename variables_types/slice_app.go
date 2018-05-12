package main

import "fmt"

func main(){
	var nums= []int{0,1,2,3,4,5,6,7,8}
	fmt.Printf("nums=:%v, len=%d, cap=%d\n",nums,len(nums),cap(nums))
	nums=append(nums,10,11,12,13)
	fmt.Printf("nums=:%v, len=%d, cap=%d\n",nums,len(nums),cap(nums))
}
