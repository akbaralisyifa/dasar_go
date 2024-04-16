package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai, pembagi int) (int, error){  // return value bisa berupa int atau error 
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0");
	}else{
		result := nilai / pembagi;
		return result, nil;
	}
}

func main() {
	// var contohError error = errors.New("Ini Error bro !!!");
	// fmt.Println(contohError.Error());

	hasil, err := Pembagi(10, 0);
	if err == nil { // kalau tidak error
		fmt.Println("Hasil", hasil)
	}else {
		fmt.Println("Error", err.Error())
	}
}