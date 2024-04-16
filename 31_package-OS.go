package main

import (
	"fmt"
	"os"
)

// package os lengkap di : https://pkg.go.dev/os

func main() {
	// method 1 : argument
	Args := os.Args
	fmt.Println("Package", Args)

	// method 2 : hostname // mengembalikan 2 param berupa string, error
	name, err := os.Hostname();

	if err == nil {
		fmt.Println("Hostname", name)
	}else{
		fmt.Println("Error", err)
	}
}