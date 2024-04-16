package main

import "fmt";

func interfaceKosong(i int) interface{}{ // agar tidak perlu menginisialisasi redur value nya berbentuk apa
	if i == 1 {
		return 1
	}else{
		return "herman"
	}	
}

func main(){
	var data interface{} = interfaceKosong(1);

	fmt.Println(data)
}