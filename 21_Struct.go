package main

import "fmt"

type Customer struct { // key nya = type || nama struct = Customer || lalu kata kunci = struct
	Name, Address string
	Age           int
	Merried       bool
}

// Stract function or strat method
func (customer Customer) sayHello(){
	fmt.Println("Hallo Bro, ", customer.Name)
}

func main() {
	eko := Customer{
		Name:    "Eko",
		Address: "Cirebon",
		Age:     30,
		Merried: true,
	}

	// or
	budi := Customer{"Budi", "Cirebon", 30, true}

	fmt.Println(eko);
	fmt.Println("saya :", budi)

	
	// Struct Method
	rully := Customer{Name: "Rully"};
	rully.sayHello();

}