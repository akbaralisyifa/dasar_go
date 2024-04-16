package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7)) // fungsi untuk membulankan ke atas / ke bawah sesuai yang paling deket
	fmt.Println(math.Round(1.3)) // ini juga
	fmt.Println(math.Floor(1.3)) // fungsi untuk membulan kan ke bawah
	fmt.Println(math.Ceil(1.3)) // fungsi untuk membulatkan ke atas

	fmt.Println(math.Max(1, 2)) // mengembalikan  nilai yang paling besar
	fmt.Println(math.Min(1, 2)) // mengembalikan nilai yang paling kecil 
}