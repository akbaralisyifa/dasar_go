package main

import "fmt"

func main(){
	var mount = [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"};
	var slice = mount[4:7];
	fmt.Println(slice);
	// output : [Mei Juni Juli]

	fmt.Println(len(slice)); // mengetahui panjang  = 3
	fmt.Println(cap(slice));  // mengetahu kapaitas = 8

	// append (menambahkan data)
	var slice2 = mount[10:];
	fmt.Println(slice2);

	var appendData = append(slice2, "Bulan Ramadhan");
	fmt.Println(appendData);

	// make
	var newSlice = make([]string, 2, 5); // 2 = panjang, 5 = kapasitas

	newSlice[0] = "bambang";
	newSlice[1] = "herman";

	fmt.Println(newSlice);

	// copy make
	copySlice := make([]string, len(newSlice), cap(newSlice)); // len = panjang, cap = kapasitas
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}