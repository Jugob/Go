package main
import "fmt"

func main(){
	var myIntArray= [5]int{4,5,6,7,8}
	fmt.Println(myIntArray)
	var avg float64=0
	var total int =0
	for i:=0;i<len(myIntArray);i++{
		total+=i
	}
	avg=float64(total)/float64(len(myIntArray))
	fmt.Println("avg=",avg)
}

