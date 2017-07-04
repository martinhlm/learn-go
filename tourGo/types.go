package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func structsFunc() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Struct fields can be accessed through a struct pointer.
	p := &v
	p.X = 1e9
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		r  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, r, v2, v3)	
}

func arraysFunc() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slicesFunc() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	// A slice does not store any data, it just describes 
	// a section of an underlying array
	// Changind the elements of a slice modifies the corresponding
	// elements of its underlying array
	names := [4]string {
		"John","Paul","Ringo", "George",
	}
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a,b)
	fmt.Println(names)
}

func more_types() {
	fmt.Println("pointers")
	pointers()
	fmt.Println(Vertex{1, 2})
	structsFunc()
	arraysFunc()
	slicesFunc()
}
