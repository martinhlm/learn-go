package main

import (
	"fmt"
	"io" // specifies the io.Reader interface, which represents the read
	"math"
	"time"
	// end of a stream of data
	"strings"
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

type Person struct {
	Name string
	Age  int
}

type MyError struct {
	When time.Time
	What string
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

	a = f  // a MyFloat implements Abser
	a = &v // a *Mertex implements Abser

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

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringersFunc() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zhap Beeble", 9001}
	fmt.Println(a, z)
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// functions often return an error value, and calling code should handle errors
// by testing whether the errors equals nil
func errorsFunc() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func readersFunc() {
	r := strings.NewReader("Hello, reader")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func interfaces() {
	interfacesFunc()
	typeAssertions()
	stringersFunc()
	errorsFunc()
	readersFunc()
}
