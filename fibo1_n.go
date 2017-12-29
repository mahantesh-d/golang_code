package main

import "fmt"

func main(){
	var i int
	a := 0
	b := 1
	fib := 0
	fmt.Print("Enter the fibo series number: ")
	_, err := fmt.Scanf("%d", &i)
	fmt.Print(a)
	if err != nil{
		fmt.Println("Enter the integer")
	}
	for j:=1;j<i;j++{
		fib = a + b
		fmt.Print("\t",fib)
		b = a
		a = fib
	}

}
