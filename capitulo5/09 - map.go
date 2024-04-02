package main

import "fmt"

func main() {

	my_map := make(map[string]int)

	my_map["Carlos"] = 9
	my_map["Ana"] = 10
	my_map["Luke"] = 8
	my_map["Yoshi"] = 8
	my_map["Amora"] = 7

	fmt.Println(my_map)
	fmt.Println(len(my_map))

	delete(my_map, "Amora")

	fmt.Println(my_map)
	fmt.Println(len(my_map))

}
