package main

import (
	"fmt"
	"strings"
)

func wordCount(text string) map[string]int {
	counts := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		word = strings.ToLower(strings.Trim(word, ".,!?;:\"'"))
		if word == "" {
			continue
		}
		counts[word]++
	}
	return counts
}

func main() {
	text := "Go is fast. Go is simple, and Go is fun! Is it fast? Yes."
	counts := wordCount(text)

	order := []string{"go", "is", "fast", "simple", "and", "fun", "it", "yes"}
	for _, w := range order {
		if c, ok := counts[w]; ok {
			fmt.Printf("%-8s %d\n", w, c)
		}
	}
	fmt.Println("total unique words:", len(counts))
}
