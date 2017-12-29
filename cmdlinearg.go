package main

import (
	"os"
	"fmt"
)

func main(){
	a := os.Args[1]
	fmt.Print("Hello ",a)
}
