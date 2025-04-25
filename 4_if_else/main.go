package main

import "fmt"

func main(){
	var x int 

	fmt.Print("Enter Number: ")
	_, err := fmt.Scan(&x)

	if err != nil {
		fmt.Println("Error reading input:" , err)
		return
	}


	if x < 10 {
		fmt.Println("less than 10")
	}else if x < 20 {
		fmt.Println("less than 20")
	}else if x == 20 {
		fmt.Println("Equal to 20")
	} else {
		fmt.Println("Greater than 20")
	}
}
