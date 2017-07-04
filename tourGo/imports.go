package main

import ( 
	"fmt" 
	"math"
)

func add(x, y int) int {
	return 	x + y
}

func main() {
	fmt.Println("Now you have %g problems.", math.Sqrt(7))
	fmt.Println("add: %g", add(42, 13))
}
