package main

import "fmt";

func main() {
	var (
		a = 10;
	 	b = 10;
	)

	var tambah = a + b;
	var kurang = a - b;
	var kali = a * b;
	var bagi = a / b;
	var sisa = a % b;

	fmt.Println(tambah);
	fmt.Println(kurang);
	fmt.Println(kali);
	fmt.Println(bagi);
	fmt.Println(sisa);

	// Augmented Assignment
	a = a + 10;
	// or
	a += 10;


	// operation perbandingan
	const aValue = "bare";
	const bValue = "bare";

	fmt.Println(aValue == bValue);

	// boolear operation 
	const uts = 80;
	const uas = 80;

	const lulus  = uts >= 80 && uas >= 80;

	fmt.Println(lulus)
}