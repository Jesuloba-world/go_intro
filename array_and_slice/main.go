package main

import "fmt"

func main() {
	a1 := [...]int{2, 4, 6, 4, 8} //array
	fmt.Printf("%#v\n", a1)

	slice := []int{2, 4, 6, 4, 8} //slice
	fmt.Printf("%#v\n", slice)

	type name []string
	var friends = name{"John", "Paul"}
	friends[0] = "Matt"
	fmt.Printf("%#v\n", friends)
}
