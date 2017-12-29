package main

import "fmt"

func main(){
	var n int
	var reversenum int
	fmt.Print("Enter the Number to check palindrome: ")
	_, err := fmt.Scanf("%d", &n)
	tmp := n
	if err != nil{
		fmt.Print("Enter the Integer")
	}
	for n != 0{
		reversenum = reversenum * 10 + n % 10
		n = n/10
	}
	fmt.Print("\n Really Number is: ", tmp)
	fmt.Print("\n Reversed Number is: ", reversenum)
	if reversenum == tmp{
		fmt.Print("\n Number is Palindrome")
	}else{
		fmt.Print("\n Number is not Palindrome")
	}
}
