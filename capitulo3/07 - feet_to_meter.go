package main

import "fmt"

func main() {

	const feet float64 = 0.3048

	fmt.Println("Input a value in meters: ")

	var meters float64

	fmt.Scanf("%f", &meters)

	result := fmt.Sprintf("%.2f %s %.2f %s", meters, "meters is equivalent to", (feet * meters), "feet(s)")

	fmt.Printf(result)

}
