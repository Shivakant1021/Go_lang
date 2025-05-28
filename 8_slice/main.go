package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7, 8)
	numbers = append(numbers, 6, 7, 8)
	fmt.Println("numbers:", numbers)
	// fmt.Printf("Number has data type : %T /n", numbers)
	fmt.Println("Length:", len(numbers))
	name := []string{}
	fmt.Println("name:", name)
	fmt.Println("Capacity:", cap(numbers))  // Capacity is the maximum number of elements that can be stored in the slice without reallocating memory
	person := []int{}
	fmt.Println("person:", person)
	fmt.Println("Capacity:", cap(person)) 
}