package main

import (
	"fmt"
	"math"
)

type Mertex struct {
	X,Y float64
}

type MyFloat float64

type Abser interface {
	Abs() float64
}

/* 
	Go does not have classes. However, you can define methods on types.
	A method is a function with a special 'receiver' argument.
	The receiver appears in its own argument list between the func keyword
	and the method name
*/
func(v *Mertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methods() {
	v := Mertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

/*
	You can declare a method on non-struct types.
	You can only declare a method with a receiver whose type is defined in the
	same package as the method.
*/
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
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
}

func methodsInterfaces() {
	methods()
	interfacesFunc()
}
