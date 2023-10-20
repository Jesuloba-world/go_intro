package main

import (
	"fmt"
)

func main() {
	// var n1 = 21
	s1 := 34043
	f1 := "3.143"

	var convertedS1 = fmt.Sprintf("%d", s1)

	fmt.Printf("%T\n", convertedS1)

	// var convertedF1 = strconv.parseFloat(f1)

	_ = f1
}
