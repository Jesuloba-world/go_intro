package main

import "fmt"

const secondsInHour = 60 * 60 * 24

func main() {
	fmt.Println("Hello Go world!")
	distance := 60.8
	fmt.Printf("The distance in miles is %f \n", distance*0.621)
}
