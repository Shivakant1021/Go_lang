package main

import "fmt"

func Sum(a *[4]float32) (sum float32) {
	for _, v := range *a {
		sum += v
	}
	return
}

func main() {
	array := [...]float32{7.4, 8.5, 6.3, 8.4}
	x := Sum(&array)
	fmt.Println(x)
}
