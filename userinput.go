package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	reader.Scan()
	text := reader.Text()
	fmt.Println("Hello ", text)
}
