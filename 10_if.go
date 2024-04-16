package main

import "fmt";

func main(){
	name := "tai";

	if name == "barre"{
		fmt.Println("Hello " + name);
	}else if  name == "yanto" {
		fmt.Println("Halo " + name)
	}else {
		fmt.Println("Hi " + name + ", kenalan YOK ?!")
	}

	// sort statment
	// cara 1
	length := len(name);
	if length > 5 {
		fmt.Println("Nama terlalu panjang");
	}else{
		fmt.Println("Panjangnya OK");
	}

	// cara 2
	if length:= len(name); length > 5 {
		fmt.Println("Terlalu panjang");
	}else{
		fmt.Println("Panjangnya OK");
	}
}