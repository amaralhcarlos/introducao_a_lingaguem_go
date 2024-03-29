package main

import "fmt"

func main() {

	fmt.Println("Digite um numero: ")

	var number float64
	fmt.Scanf("%f", &number)

	result := fmt.Sprintf("%s%.2f%s%.2f", "O Dobro de ", number, " eh ", (number * 2))
	fmt.Printf(result)

}
