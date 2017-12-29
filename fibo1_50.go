package main

import "fmt"

func main(){
	a := 0
	b := 1
	fib := 0
	fmt.Print(a)
	for i:=1;i<50;i++{
		fib = a + b
		fmt.Print("\t",fib)
		b = a
		a = fib
	}
}
