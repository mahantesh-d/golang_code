package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the number between 1-7: ")
	reader.Scan()
	text := reader.Text()
	if text == "1" {
		fmt.Print("Its Mon")
	}else if text == "2" {
		fmt.Print("Its Tue")
	}else if text == "3" {
		fmt.Print("Its Wed")
	}else if text == "4" {
		fmt.Print("Its Thur")
	}else if text == "5" {
		fmt.Print("Its Fri")
	}else if text == "6" {
		fmt.Print("Its sat")
	}else if text == "7" {
		fmt.Print("Its sun")
	}else {
		fmt.Print("Enter between 1-7")
	}
}
