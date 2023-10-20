package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func vfunc(a ...int) {
	fmt.Println(a)
}

func vfunc2(a ...float64) (float64, float64) {
	sum := float64(0)
	product := float64(1)
	for _, v := range a {
		sum += v
		product *= v
	}

	return sum, product
}

func main() {
	mySum := sum(5, 7)

	fmt.Println(mySum)

	vfunc(1, 2, 3, 45, 66)

	sum, product := vfunc2(1.2, 3.4, 6, 8)
	fmt.Println(sum, product)
}
