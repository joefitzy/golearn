// Variables and primitive Types
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 39487 // variable and type declaration
	intNum++
	fmt.Println(intNum)

	var uIntNum uint = 23432
	fmt.Println(uIntNum)

	var floatNum float64 = 1235678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	result := floatNum32 + float32(intNum32)
	fmt.Println(result)

	var myString string = "Hello World"
	myString = `Hello
	World`
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBool bool
	fmt.Println(myBool)

	var intNum3 rune
	fmt.Println(intNum3)

	// type and variable inference
	myVar := "text"
	fmt.Println(myVar)

	const myConst = "some constant value"
	fmt.Println(myConst)

	const pi = 3.1415
	fmt.Println(pi)
}
