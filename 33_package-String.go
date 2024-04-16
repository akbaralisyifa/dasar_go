package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello World", "World")) // untuk mengecek apakah World ada di dalam kata "Hellow World"; (true or false)
	fmt.Println(strings.Split("Hello World", " ")) // untuk membuat kata yang di dalam nya menjadi sslice
	fmt.Println(strings.ToLower("Hello World")) // mengubah kata menjadi kecil
	fmt.Println(strings.ToUpper("Hello World")) // mengubah kata menjadi besar
	fmt.Println(strings.Trim("   Hello World   ", " ")) // menghilangkan spasi di awal dan akhir (untuk mengahus data kiri dan kanan nya)
	fmt.Println(strings.ReplaceAll("eko pergi sama eko", "eko", "budi")) // untuk mengubah kata dari apa menjadi apa // kalau replace saja berati hanya 1
}