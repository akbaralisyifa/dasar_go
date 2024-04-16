package main

import "fmt";

func main() {
	// BAGIAN VARIABLE
	var name string;
	name = "Barre";
	fmt.Println(name);

	// atau
	var umur = 22;
	fmt.Println(umur)

	// atau 
	alamat := "Jl. Jend. Gatot Subroto";
	fmt.Println(alamat);

	// atau
	var (
		firstName = "muhammad";
		lastName = "Ali";
	);

	fmt.Println(firstName, lastName);

	// KONVERSI NILAI 
	var nilai32 int32 = 23;
	var nilai64 int64 = int64(nilai32);

	fmt.Println(nilai64);


	// BAGIAN CONSTANT
	const tanggal = 23;  // tidak akan error kalau di deklarasikan
	fmt.Println(tanggal)

	// atau
	const (
		bulan = 12;
		tahun = 2022;
	)
	fmt.Println(bulan, tahun);
}