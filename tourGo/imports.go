package main

import (
	"fmt"
	"math"
	//"math"
)

func vars() {
	var c, python, java bool
	cpp, py, jav := true, false, "no!"
	const Pi = 3.14

	fmt.Println(c, python, java, cpp, py, jav, Pi)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	vars()
	fmt.Println("Now you have %g problems.", math.Sqrt(7))
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(19))
	flow()

	more_types()
	methodsInterfaces()
	interfaces()
	concurrency()
}
