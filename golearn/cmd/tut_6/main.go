// Structs
package main

import "fmt"

type engine interface {
	milesLeft() uint8
}

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

// method on gasEngine
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type electricEngine struct {
	mpkwh     uint8
	kwh       uint8
	ownerInfo owner
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type owner struct {
	name string
}

// Using interface as function arg type
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	// var myEngine gasEngine = gasEngine{23, 12, owner{"Joe"}}
	myEngine := gasEngine{23, 12, owner{"Joe"}}
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo)

	// Anonymous structs
	var engine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}

	fmt.Println(engine2)

	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	myElectricEngine := electricEngine{25, 15, owner{"Joe"}}
	canMakeIt(myElectricEngine, 50)
}
