package main

import "fmt"

func main() {
	// 1.assigning variables
	var name string = "Golang"
	fmt.Println(name)

	// 2.assigning variables
	var class = "three"
	fmt.Println(class)

	// 3. direct assigning method
	a := 33
	fmt.Println(a)

	// 4. with const

	const isAdult = true
	fmt.Println(isAdult)
	// bonus tip for me
	if isAdult == true {
		fmt.Println("yeh bro he is adult what your Problem ?")
	}
}

/*
    Golang
	three
	33
	true
	yeh bro he is adult what your Problem ?
    

*/