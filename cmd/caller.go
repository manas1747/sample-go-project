package main

import (
	"fmt"
	"sample-go-project/pkg/printUtils"
	"sample-go-project/pkg/simpleinterest"
	"sample-go-project/pkg/utils"
)

func init() {
	fmt.Println("main::init()")
}

func main() {
	fmt.Println("Hello World")

	a, b := 11, 22
	fmt.Println("a =", a, "b =", b)

	a, b = 33, 22
	fmt.Println("a =", a, "b =", b)

	fmt.Println(simpleinterest.Calculate(11.00, 22.00, 33.00))
	fmt.Println(utils.Name)

	printUtils.PrintAstericks()

	switch num := utils.GetNumber(); {
	case num < 50:
		fmt.Printf("%d is less than %d\n", num, 50)
	case num < 100:
		fmt.Printf("%d is less than %d\n", num, 100)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than %d\n", num, 200)
	}

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
}
