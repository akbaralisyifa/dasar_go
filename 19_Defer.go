package main

import "fmt"

/**
adalah fungsi yang bisa kita jadwalkan untuk di eksekusi
setelah sebuah fungsi di eksekusi ( klo di js kaya Promise )
*/

func logging(){
	fmt.Println("Selesai memanggil function");
};

func runApplication(){
	defer logging(); // gunakan defer di atas, agar kalau error tetap di runing
	fmt.Println("run Application");
}

func main(){
	runApplication();
}