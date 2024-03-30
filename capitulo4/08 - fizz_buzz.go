package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {

		var result string

		if i%3 == 0 && i%5 == 0 {

			result = "FizzBuzz"

		} else if i%3 == 0 {

			result = "Fizz"

		} else if i%5 == 0 {

			result = "Buzz"

		} else {

			result = fmt.Sprintf("%d", i)

		}

		fmt.Println(result)

	}
}
