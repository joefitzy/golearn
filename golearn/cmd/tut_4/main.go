// Arrays, Slices, Maps, Ranges
package main

import "fmt"

func main() {

	// Static Arrays have static lengths
	// [...] infers static length

	var intArr [3]int32
	fmt.Println(intArr)

	intArr1 := [3]int32{1, 2, 3}
	fmt.Println(intArr1)

	intArr2 := [...]int32{1, 2, 3}
	fmt.Println(intArr2)

	// Slices
	var intSlice []int = []int{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	intSlice2 := []int{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice, intSlice2)

	// intSlice3 := make([]int, 3, 8) // make slice with capacity

	// Maps

	var myMap = map[string]uint{"Adam": 52, "Sarah": 32}
	fmt.Println(myMap["Adam"])

	delete(myMap, "sarah")

	var age, ok = myMap["Benny"]
	if ok {
		fmt.Printf("Age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// Range based For loop

	for name, age := range myMap {
		fmt.Printf("Name: %v, Age: %v", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v", i, v)
	}
	var i int
	for i < 10 {
		fmt.Println("something")
		i++
	}
}
