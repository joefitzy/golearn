// Strings, Runes, and Builders
package main

import (
	"fmt"
	"strings"
)

func main() {
	// utf-8 and bytes
	myString := "résumè"
	index := myString[1]
	fmt.Printf("%v, %T\n", index, index)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	// runes
	myRunes := []rune("résumé")
	var index2 = myRunes[1]
	fmt.Printf("%v is %T\n", index2, index2)
	for i, v := range myRunes {
		fmt.Println(i, v)
	}

	// String Builder
	strSlice := []string{"g", "o", "l", "a", "n", "g"}
	var strBuilder strings.Builder
	for _, v := range strSlice {
		strBuilder.WriteString(v)
	}
	fmt.Printf("\n%v", strBuilder.String())

}
