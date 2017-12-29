package main

import "fmt"

func main(){
	var n int
	var digits int
	sumOfDigit := 0
	fmt.Print("Enter the number, for summation of digits: ")
	fmt.Scanf("%d", &n)
	for n != 0 {
		digits =  n % 10
		sumOfDigit = sumOfDigit + digits
		n = n/10

	}
	fmt.Print(sumOfDigit)
}
