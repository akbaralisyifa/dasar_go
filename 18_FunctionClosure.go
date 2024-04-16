package main

import "fmt";

func main(){
	counter := 0;

	increment := func(){
		counter++ // hati2 jangan gunakan variable yang sama di luar function ini
	}

	fmt.Println(counter);
}