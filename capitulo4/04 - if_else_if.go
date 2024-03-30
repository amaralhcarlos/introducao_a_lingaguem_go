package main

import "fmt"

func main() {

	i := 2

	if i%2 == 0 {
		fmt.Println("I é divisivel por 2")
	} else if i%3 == 0 {
		fmt.Println("I é divisivel por 3")
	} else if i%4 == 0 {
		fmt.Println("I é divisivel por 4")
	}
}
