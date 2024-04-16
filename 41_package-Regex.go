package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z])o`);

	fmt.Println(regex.MatchString("eko")); // true
	fmt.Println(regex.MatchString("eLo")); // false

	search := regex.FindAllString("eko, elo, eti, samsul, bambangs", -1); // masukkan data nya, lalu berapa data yang ingin di cari ( -1 = semua )
	fmt.Println(search);
}