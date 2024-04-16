package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New() // untuk membuat list 
	data.PushBack("Bambang") // menambahkan data yang paling ujung
	data.PushBack("Ba") // menambahkan data yang paling ujung
	data.PushBack("Bang") // menambahkan data yang paling ujung

	fmt.Println(data.Front().Value) // untuk mengambil data yang paling depan
	fmt.Println(data.Back().Value) // untuk mengambil data yang paling  belakang

	// mapping dari depan ke belakang 
	for e := data.Front(); e != nil; e = e.Next(){
		fmt.Println( "data Mapping :", e.Value)
	}

	// mappiing dari belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev(){
		fmt.Println( "data Mapping :", e.Value)
	}

}