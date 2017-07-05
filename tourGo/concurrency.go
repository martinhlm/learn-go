package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutines() {
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum	//send sum to c
}

func channelsFunc() {
	s := []int{7,2,8,-9}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <- c, <- c 		// receive from c

	fmt.Println(x, y, x+y)
}

func concurrency() {
	goroutines()
	channelsFunc()
}
