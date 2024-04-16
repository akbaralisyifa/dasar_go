package main

import "fmt";

func main() {
	name := "bambang";

	switch name {
	case "bambang":
		fmt.Println("Hello bambang");
	case "barre":
		fmt.Println("Hello barre");
	default:
		fmt.Println("hi boleh kenalan ga?")
	}

	// sort statment
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu panjang");
	case false:
		fmt.Println("Panjangnya OK");
	}

	// or
	var panjang = len(name);
	switch{
	case panjang > 15:
		fmt.Println("Terlalu panjang");
	case panjang  > 5:
		fmt.Println("nama masih panjang");
	default:
		fmt.Println("Panjangnya OK");
	}
}