package main

import "fmt";

type filter func(string) string // untuk menghemat tempat menggunakan type

func sayHelloWithFilter(name string, filter filter ){ // func(string) string
	nameFiltered := filter(name);
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string{
	if name == "anjing"{
		return "..."
	}else{
		return name;
	}
}

func main() {
	sayHelloWithFilter("bambang", spamFilter);
	sayHelloWithFilter("anjing", spamFilter);

}