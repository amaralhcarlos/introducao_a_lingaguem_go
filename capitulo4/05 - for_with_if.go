package main

import "fmt"

func main() {

	var result string

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			result = fmt.Sprintf("%d%s", i, " é par")
		} else {
			result = fmt.Sprintf("%d%s", i, " é impar")
		}

		fmt.Println(result)
	}
}
