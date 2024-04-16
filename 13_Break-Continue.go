package main

import "fmt"

func main() {

	// BREAK
	for i := 0; i < 10; i++ {
		fmt.Println("break ke = ", i)
		if i == 5 {
			break
		}
	}

	// CONTINUE
	for i:= 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}

		fmt.Println("continue ke = ", i)
	}
}