package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	} else {
		return a / b, nil
	}
}

func main() {
	ans, _ := divide(10, 3)
	fmt.Println("The answer is", ans)

}
