package main

import "fmt"

func add(a, b int) int{
	return a + b
}

func main() {
	defer fmt.Println("Start of main function")
	fmt.Println("middle of main function")
	result := add(3, 4)
	defer fmt.Println("Result of addition:", result)  
	fmt.Println("End of main function")
}

// defer keyword defers the execution of this line until the surrounding function returns and it go in LIFO order