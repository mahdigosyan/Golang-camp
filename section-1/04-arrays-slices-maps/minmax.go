package main

import (
	"errors"
	"fmt"
)

func minMax(numbers []int) (int, int, error) {
	if len(numbers) == 0 {
		return 0, 0, errors.New("slice is empty")
	}

	min := numbers[0]
	max := numbers[0]

	for _, n := range numbers[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max, nil
}

func main() {
	numbers := []int{42, -7, 15, 88, 3, 0, -19, 50}

	min, max, err := minMax(numbers)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("len: %d, cap: %d\n", len(numbers), cap(numbers))
	fmt.Printf("min: %d, max: %d\n", min, max)

	empty := []int{}
	_, _, err = minMax(empty)
	fmt.Println("empty slice error:", err)
}
