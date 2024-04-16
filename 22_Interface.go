package main

import "fmt"

// jadi interface itu di gunakan untuk penjelasan sebuah method atau fungsi untuk stract
type HasName interface {
	getName() string
}


// implementasi interface dapatkan name nya;
func sayHello( hasname HasName){
	fmt.Println("hello", hasname.getName())
}

type Person struct {
	Name string;
}

// struct animal
type Animal struct {
	Name string;
}

// inisialisasi struct nya
func (person Person) getName() string {
	return person.Name;
}

func (animal Animal) getName() string {
	return animal.Name;
}


func main() {
	eko := Person{Name:  "Eko"};

	kucing := Animal{Name: "Cater" }

	sayHello(kucing)

	sayHello(eko)
}