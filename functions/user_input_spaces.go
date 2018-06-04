package main

import (
	"fmt"
	"bufio"
	"os"
//	"strings"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string")
	line,_:=reader.ReadString('6')
	//ReadString takes a "delimiter" as input. It keeps reading the input
	//until it gets that delimiter as input. Then it stops.
	fmt.Println(line)
}
