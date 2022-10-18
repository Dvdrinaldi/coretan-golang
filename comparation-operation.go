package main

import "fmt"

func main() {
	var name1 = "David"
	var name2 = "David"

	var result bool = name1 == name2
	fmt.Println(result)

	var val1 = 100
	var val2 = 200

	fmt.Println(val1 > val2)
	fmt.Println(val1 < val2)
	fmt.Println(val1 == val2)
	fmt.Println(val1 != val2)
}
