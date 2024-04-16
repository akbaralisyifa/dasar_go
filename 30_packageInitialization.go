package main

import (
	"fmt"
	"latihan1/database"
)

func main() {
	result := database.GetDatabse();
	fmt.Println(result)
}