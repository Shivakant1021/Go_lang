package main

import (
	"fmt"
	"reflect"
)

func main() {
	// simple values
	fmt.Println(1)

	// String values
	fmt.Println("This is String")

	//Float values
	fmt.Println(56.5)

	//integers
	fmt.Println(6 + 8)

	//Bool values
	fmt.Println(true)

	fmt.Println(false)

	//type of
	var a = 10
	fmt.Println("type of a:", reflect.TypeOf(a))
	fmt.Println(78.9 / 3.6)
}


/*
   // Outputs
  	 1
	This is String
	56.5
	14
	true
	false
	type of a: int
	21.916666666666668

*/