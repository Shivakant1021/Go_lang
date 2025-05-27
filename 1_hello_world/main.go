package main

import "fmt"



func main() {
	fmt.Println("hello World")
	var age int
	age = 25
	fmt.Println(age)  // Print the value of the variable 'age'
	fmt.Println(&age) // Print the memory address of the variable 'age'
	name := "John Doe"
	fmt.Println("name is", name)
	fmt.Println(&name) // Print the memory address of the variable 'name'

	id := 569
	fmt.Println("id is", &id)
}
