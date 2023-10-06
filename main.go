// A "Hello Wolrd" progam in go. Print statement and current time
package main

import (
	"fmt"
	"time"
)

func greeting() string {
	return "Hello World. The Time is: " + time.Now().String()
}

// This a "doc comment" because it immediatley proceeds a top level func
// needlesslyComplexGreeting uses needlessly complex operations to return a
// pleasant, semi-useful greeting.
func NeedlesslyComplexGreeting() string {
	var now string
	now = time.Now().String()
	msg := "Hello world, the time is: "
	msg += now
	return msg
}

func simpliedComplexGreeting() string {
	now := time.Now().String()
	msg := "Hello world, the time is: " + now
	return msg
}

func main() {
	fmt.Println(greeting())
}
