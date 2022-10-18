package main

import "fmt"

func main() {
	var name string

	name = "David Rinaldi"
	fmt.Println(name)

	name = "Dav"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	var gang int8 = 5
	fmt.Println(gang)

	negara := "Indonesia"
	fmt.Println(negara)

	negara = "Malaisya"
	fmt.Println(negara)

	var (
		firstName = "David"
		lastName  = "Rinaldi"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
