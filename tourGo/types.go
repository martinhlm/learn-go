package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

type Vert struct {
	Lat, Long float64
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
	var p []int = primes[1:4]
	fmt.Println(p)

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

	// defaults
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func sliceAppend() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
		printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func rangeFunc() {
	var pow = []int{1, 2, 4, 8}

	for i,v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// you can skip the index or value by assigning to _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func mapsFunc() {
	var m map[string]Vert

	// make function returns a map of the given type, initialized and ready
	m = make(map[string]Vert)
	m["Bell labs"] = Vert{40.234, -74.33289,}
	fmt.Println(m["Bell labs"])

	var ma = map[string]Vert{
		"Bells lab": Vert{
			40.3333, -23.9999,
		},
		"google": Vert{
			98.4444, 45.2222,
		},
	}
	fmt.Println(ma["Bells lab"])
	
	// if the top level type is just a type name, you can omit it from the
	// elements of the literal
	var mapp = map[string]Vert{
		"Bells lab": { 40.3333, -23.9999 },
		"google": { 98.4444, 45.2222 },
	}
	fmt.Println(mapp["Bells lab"])
}

func mutatingMaps() {
	m := make(map[string]int)

	m["answer"] = 24
	fmt.Println("the value: ", m["answer"])
	m["answer"] = 48
	fmt.Println("the value: ", m["answer"])
	delete(m, "answer")
	fmt.Println("the value: ", m["answer"])
	v, ok := m["answer"]
	fmt.Println("the value: ", v, "present?", ok)
}

func more_types() {
	fmt.Println("pointers")
	pointers()
	fmt.Println(Vertex{1, 2})
	structsFunc()
	arraysFunc()
	slicesFunc()
	sliceAppend()
	rangeFunc()
	mapsFunc()
	mutatingMaps()
}
