package main

import "fmt"

func main() {

	slice := []int{1, 2, 3}

	slice2 := make([]int, 2)

	copy(slice2, slice)

	fmt.Println(slice)
	fmt.Println(slice2)

}
