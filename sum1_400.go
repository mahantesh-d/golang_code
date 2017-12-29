package main

import "fmt"

func main(){
	var sum int
	for i:=0; i<=400; i++ {
		sum = sum+i
	}
	fmt.Print(sum)
}
