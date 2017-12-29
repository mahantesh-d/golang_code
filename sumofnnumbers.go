package main

import (
	"fmt"
)

func main(){
	var i int
	sum := 0
	fmt.Print("Enter the number: ")
	_, err := fmt.Scanf("%d", &i)
	if err != nil{
		fmt.Print("Enter the Integer")
	}
	for j := 0; j<=i; j++ {
		sum = sum+j
	}
	fmt.Print(sum)
}
