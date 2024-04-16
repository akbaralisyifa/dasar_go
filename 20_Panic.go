package main

import "fmt"

/**
*Panic : Fungsing yang bisa kita gunakan untuk menghentikan Program,
fungsi ini di panggil ketika terjadi error pada saat program kita berjalan
( seperti TryCatch / thenCatch )

*Recover : fungsi yang dapat menangkap data Panic,
dengan recover panic akan berhenti sehingga program tetap berjalan;
*/

func endApp() {
	messesage := recover();
	if messesage != nil{
		fmt.Println("Error dengan message:", messesage)
	}

	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool){
	defer endApp();

	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("aplikasi berjalan")
}

func main(){
	runApp(false)
}