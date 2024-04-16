package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host","localhost","put your database host") // name host nya , lalu default value, dan message nya ()
	user := flag.String("User", "root", "put your user database ")
	password := flag.String("Password", "root", "Put your password database ")

	// Pars terlebih dahulu sebelum di eksekusi
	flag.Parse()

	fmt.Println(*host, *user, *password) // gunakan bintang di depan nya
}