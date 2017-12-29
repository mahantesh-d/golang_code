package main

import "fmt"

func main(){
	var n int
	fmt.Print("Enter the number to be reversed: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil{
		fmt.Print("Enter the digit")
	}
	var reversenum int
	for n != 0{
		reversenum = reversenum * 10 + n%10
		n = n/10
	}
	fmt.Print(reversenum)
}
