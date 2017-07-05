package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

type Abser interface {
	Abs() float64
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func interfacesFunc() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Mertex{3, 4}

	a = f 	// a MyFloat implements Abser
	a = &v 	// a *Mertex implements Abser

	// v is Mertex and does NOT implement Abser
	//a = v //error

	fmt.Println(a.Abs())

	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}

func interfaces() {
	interfacesFunc()
	typeAssertions()
}

