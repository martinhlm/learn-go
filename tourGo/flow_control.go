package main

import ( 
	"fmt"
	"math"
)

func forTypes() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	sum = 1
	for ; sum < 1000; {
		sum += sum
	}

	for sum < 1000 {
		sum += sum
	}

	fmt. Println(sum)
}


func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func flow() {
	forTypes()
	fmt.Println(sqrt(2))
}
