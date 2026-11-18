package main

import "fmt"

func sumTo(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("number must be greater than zero: %d", n)

	}

	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum, nil
}

func parity(number int) string {
	switch number % 2 {
	case 0:
		return "off"
	default:
		return "even"
	}
}

func main() {
	numbers := []int{1, 5, 6, -3, 8}

	for _, number := range numbers {
		result, err := sumTo(number)
		if err != nil {
			fmt.Println(number, err)
			continue
		}

		fmt.Println(
			number,
			number,
			result,
			parity(number),
		)
	}
}
