package main

import ( 
	"fmt" 
	"math"
)

func vars() {
	var c, python, java bool
	cpp, py, jav := true, false, "no!"

	fmt.Println(c, python, java, cpp, py, jav)
}

func add(x, y int) int { 
	return 	x + y
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
	fmt.Println("Now you have %g problems.", math.Sqrt(7))
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(19))
}
