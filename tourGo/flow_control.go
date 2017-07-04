package main

import ( 
	"fmt"
	"math"
	"runtime"
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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though

	return lim
}

func switchFunc() {
	fmt.Print("Go, runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}

func flow() {
	forTypes()
	fmt.Println(sqrt(2))
	fmt.Println(pow(3, 2, 10))
	switchFunc()
}
