package main

import (
	"fmt"
	"strconv"
)

//var block
var (
	firstname string = "Andre"
	lastname  string = "Mueller"
	age       int
)

// UPPERCASE variable, visible globally
var GLOBALVAR string = "global"

// var visible in package
var packagevar = "package"

func stringTest() []byte {

	//string is UTF8
	s := "Andre Mueller"
	//converting string s into a byte slice
	b := []byte(s)

	fmt.Println("Printing ascii value for first Character of the string")
	fmt.Println(s[0])

	//golang rune -> This is a type alias for int32 to be used for UTF32
	r := 'a'
	fmt.Printf("%v ,%T\n", r, r)
	return b
}

func stringByteOps() {
	var i, u int
	i = 10
	u = 3
	// binary operations
	fmt.Println("\nBinary Operations")
	fmt.Println(i & u)
	fmt.Println(i | u)
	fmt.Println(i ^ u)
	//bit shifting
	fmt.Println(i >> 3)
	fmt.Println(i << 3)
}

func main() {

	// Let compiler figure out the type
	stadt := "Münster"
	age = 53
	temperature := 36.8

	// type conversion with data loss
	var intTemperature = int(temperature)

	// String conversion, but uses unicode value for integer
	var stringAge = string(age)
	var stringAge2 = strconv.Itoa(age)

	fmt.Println(firstname, lastname, age, stadt, GLOBALVAR, packagevar, temperature, intTemperature, stringAge, stringAge2, stringTest())
	fmt.Printf("%v, %T", age, age)
	stringByteOps()
}
