package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPDavid NoKTP = "123123123123123123123"
	var marriedStatus Married = true

	fmt.Println(noKTPDavid)
	fmt.Println(marriedStatus)
}
