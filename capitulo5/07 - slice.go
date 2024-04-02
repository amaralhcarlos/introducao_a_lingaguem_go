package main

import "fmt"

func main() {

	var slice []float64 = make([]float64, 5, 10)

	slice[0] = 50
	slice[3] = 47

	fmt.Println(slice)
	fmt.Println(len(slice))

	slice = append(slice, 49)

	slice = append(slice, 88, 27, 93)

	fmt.Println(slice)
}
