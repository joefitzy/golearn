// Pointers
package main

import "fmt"

func main() {
	// memory allocation, will set to Type default
	var p *int32 = new(int32)

	fmt.Println(p, *p)
	var i int32 = 7898

	// Assign memory address to p
	p = &i
	fmt.Println(p, *p, i)

	// Assign value to the address p points too
	*p = 123
	fmt.Println(p, *p, i)

	var q **int32
	if q == nil {
		fmt.Println("unassigned pointer detected")
	}
	// Get the address of the pointer itself, not the address the pointer is pointing too
	q = &p
	printP(p)
	printP(*q)
}

func printP(p *int32) {
	fmt.Println(*p)
}
