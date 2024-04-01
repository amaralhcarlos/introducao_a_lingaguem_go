package main

import "fmt"

func main() {

	var x [5]int

	x[4] = 10

	for i := 0; i < 5; i++ {

		fmt.Printf("%d ", x[i])

	}

}
