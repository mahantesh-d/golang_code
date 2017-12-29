package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the number between 1-12: ")
	reader.Scan()
	text := reader.Text()
	switch text {
	case "1": fmt.Print("Jan")
		break
	case "2": fmt.Print("Feb")
		break
	case "3": fmt.Print("Mar")
		break
	case "4": fmt.Print("Apr")
		break
	case "5": fmt.Print("May")
		break
	case "6": fmt.Print("June")
		break
	case "7": fmt.Print("July")
		break
	case "8": fmt.Print("Aug")
		break
	case "9": fmt.Print("Sept")
		break
	case "10": fmt.Print("Oct")
		break
	case "11": fmt.Print("Nov")
		break
	case "12": fmt.Print("Dec")
		break
	default: fmt.Print("Enter between 1-12")
		break
	}
}
