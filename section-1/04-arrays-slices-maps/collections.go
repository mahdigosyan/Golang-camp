package main

import (
	"fmt"
)

func main() {
	var nilMap map[string]int

	value, ok := nilMap["missing"]
	fmt.Printf("read from nil map: value=%d, ok=%v, len=%d\n", value, ok, len(nilMap))

	readyMap := make(map[string]int)
	readyMap["go"] = 1
	readyMap["go"]++
	fmt.Printf("ready map: %v, go count=%d\n", readyMap, readyMap["go"])

	literalMap := map[string]int{"a": 1, "b": 2}
	fmt.Printf("literal map: %v\n", literalMap)

	delete(literalMap, "a")
	fmt.Printf("after delete: %v\n", literalMap)

	fmt.Printf("missing key: %d, ok=%v\n", literalMap["zzz"], func() bool {
		_, ok := literalMap["zzz"]
		return ok
	}())

}
