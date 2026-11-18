package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	englishStr := "Hello"
	persianStr := "سلام"

	fmt.Println("---برسی رشته انگلیسی---")
	analyzeString(englishStr)
	fmt.Println("---برسی رشته فارسی---")
	analyzeString(persianStr)
}

func analyzeString(s string) {
	byteCount := len(s)

	runeCount := utf8.RuneCountInString(s)

	fmt.Printf("رشته: %s\n", s)
	fmt.Printf("تعداد بایت‌ها (len): %d\n", byteCount)
	fmt.Printf("تعداد کاراکترها (RuneCount): %d\n", runeCount)
	fmt.Printf("بایت‌های تشکیل‌دهنده: %x\n", s)

}
