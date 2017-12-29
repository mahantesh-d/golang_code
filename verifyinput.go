package main

import (
	"os"
	"fmt"
)

func main(){
	a := os.Args[1:]
	if a == "Foo"{
		fmt.Print("Hi ", a)
	}else {
		fmt.Print("Dont You")
	}
}
