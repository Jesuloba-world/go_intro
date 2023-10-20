package main

import "fmt"

type names []string

func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

type Car struct {
	brand string
	price int
}

func changeCar(c Car, newbrand string, newprice int) {
	c.brand = newbrand
	c.price = newprice
}

func (c *Car) changeCar1(newbrand string, newprice int) {
	c.brand = newbrand
	c.price = newprice
}

func main() {
	friends := names{"Jesuloba", "John"}

	friends.print()

	myCar := Car{brand: "Audi", price: 49}

	fmt.Println(myCar)

	changeCar(myCar, "Benz", 31)

	fmt.Println(myCar)

	myCar.changeCar1("Benz", 31)

	fmt.Println(myCar)

}
