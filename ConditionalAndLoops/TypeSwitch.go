package main

import (
	"fmt"
)

// Type switch is a regular switch statement, but the cases check the types not
// values - the interface{} mean the passed in value can be anytype.
// check the empty interface file.
// Note: the type switch has the same syntax as type assertion i.(T), but the T
// is replaced with type
func TypeSwitch(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println("I AM AN INT", t)
	case string:
		fmt.Println("I AM A STRING", t)
	case bool:
		fmt.Println("I AM A BOOL", t)
	default:
		fmt.Println("I AM A... i don't remember...", t)
	}
}
func main() {
	TypeSwitch(9)
	TypeSwitch("Hello")
	TypeSwitch(true)
	TypeSwitch(1.2)
}
