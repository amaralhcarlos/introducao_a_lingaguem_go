package main

import "fmt"

func main() {

	fmt.Println("Input a temperature in Fahrenheit:")

	var temperature_fahrenheit float64

	fmt.Scanf("%f", &temperature_fahrenheit)

	var temperature_celsius = (temperature_fahrenheit - 32) * 5 / 9

	result := fmt.Sprintf("%.2f%s%.2f%s", temperature_fahrenheit, "F° is equal to ", temperature_celsius, "C°")

	fmt.Println(result)

}
