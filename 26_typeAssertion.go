package main

import "fmt"

func Random() interface{} {
	return "Ups" // value itu ngambil dari sini menggunakan type assertion
}

func main() {
	var result interface{} = Random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// // jika tika sesuai maka akan terjadi panic
	// var resultInt int = result.(int); // akan terjadi Panic
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value , "Istring")
	case int:
		fmt.Println("Value", value , "Int")
	default:
		fmt.Println("Unknow Type")
	}
}