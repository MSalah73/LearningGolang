package main

import (
	"fmt"
)

// If the concrte value inside the interface it self is nil, the method will be
// called wit a (nil receiver)
// Note: in some languages, this would trigger a null pointer exception, but in
// GO it is common to write methods that gracfully handle being called with a
// nil receiver ( M() in this cases)
// IMPORTANAT NOTE: interface's value that holds a nil concrete value is itself non-nil.
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T
	// This a nil interface value which holds neither value or concrete type
	// Calling M() on a nill interface is a run-time error because there os no type
	// inside the interface tuple to indicate which concrete method to call.
	// var t T -- this will result in an error -
	// T does not implement I (M method has a pointer receiver)
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v %T)n", i, i)
}
