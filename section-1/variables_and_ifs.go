package main

import "fmt"

func main() {
	var age int = 26
	name := "Mahdi"
	const bootcamp = "Golang"

	fmt.Printf("User: %s, Age: %d, Camp:%s\n", name, age, bootcamp)

	number := 42

	if number%2 == 0 {
		fmt.Printf("your number %d is odd\n", number)
	} else {
		fmt.Printf("your number %d is even\n", number)
	}

	if score := 85; score >= 90 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	var a int = 10
	var b int = 3

	var divisionResult float64 = float64(a) / float64(b)
	fmt.Printf("this is", divisionResult)
}
