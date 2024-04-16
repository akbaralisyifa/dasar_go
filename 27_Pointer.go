package main

import "fmt"

type Address struct {
	city, province, country string
}

// Pointer in Function (agar dari dari parent nya berubah tidak di duplicate)
func changeCountryToIndonesia(address *Address) {
	address.country = "Indonesia"
}

func main() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"} 
	// address2 := address1 // dia akan mengcopy data yang ada di address1 bukan menimpa 
	address2 := &address1;
	// address3 := &Address{"sulawesi", "DKI Jakarta", "Indonesia"} // dia akan membuat data baru

	address2.city = "Bandung";

	// address2 = &Address{"Surabaya", "Jawa Timur", "Indonesia"} // tanpa bintang dia akan membuat data baru dan tidak mengubah address1
	*address2 = Address{"Semarang", "Jawa Tengah", "Indonesia"} // dia akan merubah semua data di tabel yang sama ( gunakan ini)

	fmt.Println(address1)
	fmt.Println(address2)

	// or 
	address3 := new(Address);
	address3.city = "sulawesi"
	
	fmt.Println(address3)

	// Pointer in Function
	var alamat = Address{"Jakarta", "DKI Jakarta", ""};
	changeCountryToIndonesia(&alamat);
	fmt.Println(alamat);
}