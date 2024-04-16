package main

import "fmt";

func main(){
	person := map[string]string{
		"name": "barre",
		"address": "jakarta",
	}

	person["title"] = "programmer"

	fmt.Println(person);
	fmt.Println(person["name"]);
	fmt.Println(person["address"]);

	// Method lain
	books := make(map[string]string);
	books["title"] = "Sun of God";
	books["author"] = "barre";
	books["ups"] = "Wrong";

	fmt.Println(books);

	delete(books, "ups"); // untuk delete data ambil variable nya lalu, key nya apa;

	fmt.Println(books)
	fmt.Println(len(books))

}