package main

import "fmt"

func main(){
	var n int
	fmt.Print("Enter the number of digits:  ")
	fmt.Scan(&n)
	a := make([]int, n)
	fmt.Print("Enter the numbers:")
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	fmt.Println(a)
	//max := a[0]
	for i:=0;i<n-i;i++{
		if a[i]<a[i+1]{
			a[i] = a[i+1]
			fmt.Print(a[i])
		}
	}
}
