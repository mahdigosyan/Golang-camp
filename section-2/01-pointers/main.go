package main

import "fmt"

func addTenByValue(n int) {
	n = n + 10
	fmt.Printf("   [داخل تابع Value] مقدار n برابر است با: %d (آدرس محلی: %p)\n", n, &n)
}

func addTenByPointer(n *int) {
	*n = *n + 10
	fmt.Printf("   [داخل تابع Pointer] مقدار اشاره‌شده برابر است با: %d (آدرس پوینتر: %p)\n", *n, n)
}

func main() {
	fmt.Println("=== بررسی رفتار Pointer و Pass by Value در Go ===")
	scoreByValue := 50
	addTenByValue(scoreByValue)

	fmt.Printf("2. بعد از تابع Value: مقدار = %d (تغییری نکرد چون کپی بود!)\n", scoreByValue)

	fmt.Println("--------------------------------------------------")

	scoreByPointer := 50
	fmt.Printf("3. قبل از تابع Pointer: مقدار = %d، آدرس حافظه = %p\n", scoreByPointer, &scoreByPointer)
	addTenByPointer(&scoreByPointer)
	fmt.Printf("4. بعد از تابع Pointer: مقدار = %d (تغییر کرد چون مستقیماً روی حافظه اثر گذاشت!)\n", scoreByPointer)
	fmt.Println("==================================================")
}
