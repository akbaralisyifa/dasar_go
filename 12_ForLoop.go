package main

import "fmt";

func main() {
	counter := 1;

	for counter <= 10 {
		fmt.Println("perulangan ke =",counter);
		counter++
	}

	// or (For dengan statement)
	for number := 1; number <= 10; number++ {
		fmt.Println("number ke =", number)
	}

	// of ( For Range)
	slice := []string{"satu", "dua", "tiga", "empat", "lima"};

	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}
	// or
	for i, value:= range slice {  // jika index nya tidak di gunakan gan dengan "_" cth : for _, val := range slice {fmt.Println(val)}
	// fmt.Println(value)
		fmt.Println("index ", i, "=", value)
	}

	// Untuk map
	books := make(map[string]string);
	books["title"] = "Sun of God";
	books["author"] = "barre";

	for key, value := range books {
		fmt.Println(key,"=", value)
	}
}