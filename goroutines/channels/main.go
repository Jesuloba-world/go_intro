package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	c := make(chan int)

	// receiving channel
	c1 := make(<-chan string)

	// sending channel
	c2 := make(chan<- string)

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c2)

	go f1(12, c)

	n := <-c

	fmt.Println(n)
}
