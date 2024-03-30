package main

import "fmt"

func main() {

	i := 1

	switch i {
	case 0:
		fmt.Println("Zero")
		break
	case 1:
		fmt.Println("One")
		break
	case 2:
		fmt.Println("Two")
		break
	default:
		fmt.Println("Not a number")
		break
	}
}
